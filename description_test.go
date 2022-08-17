package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetDescription(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/descriptions/pulverize_description",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/descriptions/byId_pulverize_description.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	desc, err := api.GetDescription("pulverize_description")
	if err != nil {
		t.Error("error during `GetDescription`")
	}
	if desc.Name != "pulverize" {
		t.Error("error `desc.Name` is not `pulverize`")
	}
}

func TestGetDescriptionErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/descriptions/pulverize_description",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetDescription("pulverize_description")
	if err == nil {
		t.Error("error GetDescription doesn't return error`")
	}
}
