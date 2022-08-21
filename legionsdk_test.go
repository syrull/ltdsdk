package ltdsdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func LoadFixture(fileName string) any {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var fixture any
	json.Unmarshal([]byte(byteValue), &fixture)
	return fixture
}

func TestGetRequest(t *testing.T) {
	api := NewLTDSDK("test_api_key", "t[]est.com:80_invalid_port")
	m := make(map[string]string)
	err := api.getRequest("", m, "")
	if err == nil {
		t.Error("error `api.GetRequest` doesn't return error")
	}
}

func TestGetRequestJSONDecodeError(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/test",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, "{}}"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	var str string
	err := api.getRequest("test", nil, &str)
	if err == nil {
		t.Error("error `api.GetRequest` doesn't return error")
	}
}

func TestCreateAuthenticatedRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("error createAuthenticatedRequest did not panic")
		}
	}()
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	api.createAuthenticatedRequest("TEST", &url.URL{Scheme: "test", Host: "test:80_inv"})
}

func TestCachedRequest(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/abilities/byId_pulverize_melee_ability_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetAbility("pulverize_melee_ability_id")
	if err != nil {
		t.Error("error during `GetAbility`")
	}
	ability, err := api.GetAbility("pulverize_melee_ability_id")
	if err != nil {
		t.Error("error during `GetAbility`")
	}
	if ability.Name != "Pulverize" {
		t.Error("error `ability.Name` is not `Pulverize`")
	}
	if ability.Duration != 2.00000 {
		t.Error("error `ability.Duration` is not `2.00000`")
	}
}

func TestCachedRequestOnError(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/abilities/byId_pulverize_melee_ability_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetAbility("pulverize_melee_ability_id")
	if err != nil {
		t.Error("error during `GetAbility`")
	}
	// Deleting the Cache Entry and setting it to an empty byte array
	api.cache.Delete("https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id")
	api.cache.Set("https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id", []byte{})
	_, err = api.GetAbility("pulverize_melee_ability_id")
	if err == nil {
		t.Error("error `GetAbility` doesn't return error")
	}
}
