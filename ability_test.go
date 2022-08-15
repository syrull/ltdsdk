package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetAbility(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/info/abilities/byId_pulverize_melee_ability_id.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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
