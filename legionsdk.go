package ltdsdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type LegionTDSdk struct {
	client    *http.Client
	hostUrl   string
	secretKey string
}

func NewLTDSDK(secretKey string, hostUrl string) *LegionTDSdk {
	httpClient := http.DefaultClient
	return &LegionTDSdk{secretKey: secretKey, client: httpClient, hostUrl: hostUrl}
}

func (l *LegionTDSdk) createAuthenticatedRequest(method string, url *url.URL) *http.Request {
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-api-key", l.secretKey)
	return req
}

func (l *LegionTDSdk) GetRequest(endpoint string, queryString map[string]string, obj any) error {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return err
	}
	return nil
}

func toQueryString(obj interface{}) map[string]string {
	elem := reflect.ValueOf(obj).Elem()
	queryStringMap := make(map[string]string)
	for i := 0; i < elem.NumField(); i++ {
		k := elem.Type().Field(i).Tag.Get("qs")
		v := elem.Field(i).Interface()
		switch v := v.(type) {
		case int:
			if v != 0 {
				queryStringMap[k] = strconv.Itoa(v)
			}
		case string:
			if v != "" {
				queryStringMap[k] = v
			}
		}
	}
	return queryStringMap
}
