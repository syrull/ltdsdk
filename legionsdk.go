package ltdsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

type LegionTDSdk struct {
	client    *http.Client
	hostUrl   string
	secretKey string
	cache     httpcache.Cache
}

// NewLTDSDK creates a new LegionTDSdk object
func NewLTDSDK(secretKey string, hostUrl string) *LegionTDSdk {
	httpClient := http.DefaultClient
	tempDir, _ := ioutil.TempDir("", "httpcache")
	defer os.RemoveAll(tempDir)
	cache := diskcache.New(tempDir)
	return &LegionTDSdk{secretKey: secretKey, client: httpClient, hostUrl: hostUrl, cache: cache}
}

// Attaches an x-api-key header to the request
func (l *LegionTDSdk) createAuthenticatedRequest(method string, url *url.URL) *http.Request {
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("x-api-key", l.secretKey)
	return req
}

// Caching an object from a request.
func (l *LegionTDSdk) cacheObject(key string, obj any) error {
	bytes, _ := json.Marshal(obj)
	l.cache.Set(key, bytes)
	return nil
}

// getRequest performs a GET request to the API host and it serializes
// the repsonse body into a native struct
func (l *LegionTDSdk) getRequest(endpoint string, queryString map[string]string, obj any) error {
	url, err := url.Parse(fmt.Sprintf("%s%s", l.hostUrl, endpoint))
	if err != nil {
		return err
	}
	if queryString != nil {
		values := url.Query()
		for k, v := range queryString {
			values.Add(k, v)
		}
		url.RawQuery = values.Encode()
	}
	body, ok := l.cache.Get(url.String())
	if !ok {
		request := l.createAuthenticatedRequest("GET", url)
		resp, err := l.client.Do(request)
		if err != nil {
			return err
		}
		switch respCode := resp.StatusCode; {
		case respCode == 404:
			message := fmt.Errorf("entry not found")
			return message
		case respCode >= 500:
			message := fmt.Errorf("server error: %v", respCode)
			return message
		}
		if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
			return err
		}
		l.cacheObject(url.String(), obj)
		defer resp.Body.Close()
	} else {
		if err := json.Unmarshal(body, &obj); err != nil {
			return err
		}
	}
	return nil
}
