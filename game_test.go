package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetGameById(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/games/byId_314f3b.json")
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games/byId/314f3b9258f5f84b36301ac596dde8d308dcdbfbf199decbb0c1c9e0fac21321",
		httpmock.NewStringResponder(200, data))
	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	game, err := api.GetGameById("314f3b9258f5f84b36301ac596dde8d308dcdbfbf199decbb0c1c9e0fac21321")
	if err != nil {
		t.Error("error during `GetGameById`")
	}
	if game.Id != "314f3b9258f5f84b36301ac596dde8d308dcdbfbf199decbb0c1c9e0fac21321" {
		t.Error("error `game.Id` is not `314f3b9258f5f84b36301ac596dde8d308dcdbfbf199decbb0c1c9e0fac21321`")
	}
	if game.KingSpell != "lightning_hammer_activator_ability_id" {
		t.Error("error `game.KingSpell` is not `lightning_hammer_activator_ability_id`")
	}
	if len(game.PlayersData) != 8 {
		t.Error("error `len(game.PlayersData)` is not 8")
	}
}
