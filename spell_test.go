package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetSpell(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/spells/byId/allowance_powerup_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/spells/byId_allowance_powerup_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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
