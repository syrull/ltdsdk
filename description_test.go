package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetDescription(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/info/descriptions/byId_pulverize_description.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/descriptions/pulverize_description",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	desc, err := api.GetDescription("pulverize_description")
	if err != nil {
		t.Error("error during `GetDescription`")
	}
	if desc.Name != "pulverize" {
		t.Error("error `desc.Name` is not `pulverize`")
	}
}
