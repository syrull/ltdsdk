package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetAbility(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/abilities/byId_pulverize_melee_ability_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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

func TestGetAbilityErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(500, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetAbility("pulverize_melee_ability_id")
	if err == nil {
		t.Error("error `GetAbility` doesn't return error")
	}
}
