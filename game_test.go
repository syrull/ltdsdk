package ltdsdk

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

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

func TestGetGameByFakeId(t *testing.T) {
	httpmock.Activate()
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetGameById("1")
	if err == nil {
		t.Error("`GetGameById` doesn't return erorr")
	}
}

func TestGetGames(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/games/getAll.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	gameCollection, err := api.GetGames(&GameOptions{})
	if err != nil {
		t.Error("error during `GetGameById`")
	}
	if len(gameCollection.Games) != 2 {
		t.Error("error `len(*games)` != 2")
	}
}

func TestGetGamesWithErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetGames(&GameOptions{})
	if err == nil {
		t.Error("`GetGameById` doesn't return error")
	}
}

func TestGetGamesWithGameOptions(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games?dateAfter=2021-10-01+00%3A00%3A00&includeDetails=false&version=v1337",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/games/getAll.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")

	gameCollection, err := api.GetGames(&GameOptions{
		Version:   "v1337",
		AfterDate: "2021-10-01 00:00:00",
	})
	if err != nil {
		t.Error("error during `GetGames`")
	}
	if len(gameCollection.Games) != 2 {
		t.Error("error `len(games)` != 2")
	}
}

func TestExportToJson(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/games",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/games/getAll.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	gameCollection, err := api.GetGames(&GameOptions{})
	if err != nil {
		t.Error("error during `GetGames`")
	}
	file, err := ioutil.TempFile("", "tmp_export.json")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	gameCollection.ExportToJson(file.Name())
	if _, err := os.Stat(file.Name()); errors.Is(err, os.ErrNotExist) {
		t.Error("error `ExportToJson` doesn't export the file")
	}
}
