package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetSpell(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/info/spells/byId_allowance_powerup_id.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/spells/byId/allowance_powerup_id",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	spell, err := api.GetSpell("allowance_powerup_id")
	if err != nil {
		t.Error("error during `GetSpell`")
	}
	if spell.Name != "Allowance" {
		t.Error("error `spell.Name` is not `Allowance`")
	}
	if spell.Tooltip != "+100 gold" {
		t.Error("error `spell.Tooltip` is not `+100 gold`")
	}
}
