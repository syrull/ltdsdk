package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetPlayerByName(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/players/byName_syll.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byName/syll",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	player, err := api.GetPlayerByName("syll")
	if err != nil {
		t.Error("error during `GetPlayerByName`")
	}
	if player.Name != "syll" {
		t.Error("error `player.Name` is not `syll`")
	}
	if player.ID != "42A9C67482E71FEA" {
		t.Error("error `player.ID` is not `42A9C67482E71FEA`")
	}
}

func TestGetPlayerById(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/players/byId_42A9C67482E71FEA.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byId/42A9C67482E71FEA",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	player, err := api.GetPlayerById("42A9C67482E71FEA")
	if err != nil {
		t.Error("error during `GetPlayerByName`")
	}
	if player.Name != "syll" {
		t.Error("error `player.Name` is not `syll`")
	}
	if player.ID != "42A9C67482E71FEA" {
		t.Error("error `player.ID` is not `42A9C67482E71FEA`")
	}
}
