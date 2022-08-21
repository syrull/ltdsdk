package ltdsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LegionTDSdk struct {
	client    *http.Client
	hostUrl   string
	secretKey string
}

// NewLTDSDK creates a new LegionTDSdk object
func NewLTDSDK(secretKey string, hostUrl string) *LegionTDSdk {
	httpClient := http.DefaultClient
	return &LegionTDSdk{secretKey: secretKey, client: httpClient, hostUrl: hostUrl}
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
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return err
	}
	return nil
}
