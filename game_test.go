package ltdsdk

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestGetGameById(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games/byId/314f3b9258f5f84b36301ac596dde8d308dcdbfbf199decbb0c1c9e0fac21321",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/games/byId_314f3b.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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

func TestGetGames(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games/",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/games/getAll.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	games, err := api.GetGames(&GameOptions{})
	if err != nil {
		t.Error("error during `GetGameById`")
	}
	if len(games) != 2 {
		t.Error("error `len(*games)` != 2")
	}

}

func TestGameOptionsToQueryString(t *testing.T) {
	now := time.Now().Format("2000-01-01 00:00:00")
	gameOpts := &GameOptions{
		Version:        "",
		Limit:          50,
		Offset:         0,
		SortBy:         "date",
		SortDirection:  1,
		AfterDate:      now,
		BeforeDate:     now,
		IncludeDetails: false,
		QueueType:      "",
	}
	qs := toQueryString(gameOpts)
	if qs["version"] != "" {
		t.Error("error `qs['version']` != ''")
	}
	if qs["limit"] != "50" {
		t.Error("error `qs['limit']` != 50")
	}
	if qs["limit"] != "50" {
		t.Error("error `qs['limit']` != 50")
	}
	if qs["afterDate"] != now {
		t.Errorf("error `qs['afterDate']` != %s", now)
	}
}
