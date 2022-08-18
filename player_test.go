package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetPlayerByName(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byName/syll",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/players/byName_syll.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	player, err := api.GetPlayerByName("syll")
	if err != nil {
		t.Error("error during `GetPlayerByName`")
	}
	if player.Name != "syll" {
		t.Error("error `player.Name` is not `syll`")
	}
	if player.Id != "42A9C67482E71FEA" {
		t.Error("error `player.ID` is not `42A9C67482E71FEA`")
	}
}

func TestGetPlayerByNameErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byName/syll",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetPlayerByName("syll")
	if err == nil {
		t.Error("error `GetPlayerByName` doesn't return erorr")
	}
}

func TestGetPlayerById(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byId/42A9C67482E71FEA",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/players/byId_42A9C67482E71FEA.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	player, err := api.GetPlayerById("42A9C67482E71FEA")
	if err != nil {
		t.Error("error during `GetPlayerByName`")
	}
	if player.Name != "syll" {
		t.Error("error `player.Name` is not `syll`")
	}
	if player.Id != "42A9C67482E71FEA" {
		t.Error("error `player.ID` is not `42A9C67482E71FEA`")
	}
}

func TestGetPlayerByIdFakeId(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byId/fake_id",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetPlayerById("fake_id")
	if err == nil {
		t.Error("error `GetPlayerByName` doesn't return error")
	}
}

func TestGetPlayerMatchHistory(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/matchHistory/42A9C67482E71FEA",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/players/matchHistory_42A9C67482E71FEA.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	games, err := api.GetPlayerMatchHistory("42A9C67482E71FEA", &MatchHistoryOptions{})
	if err != nil {
		t.Error("error during `GetPlayerMatchHistory`")
	}
	// The default is 10, this is just for perfomance reasons
	if len(games) != 2 {
		t.Error("error `len(*games)` != 2")
	}
}

func TestGetPlayerMatchHistoryWithMatchHistoryOptions(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/matchHistory/42A9C67482E71FEA?countResults=false&limit=2",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/players/matchHistory_42A9C67482E71FEA.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	// The limit is 3 but since that we are getting a response by the mocked responder
	// with a status code 200, that means that the query param is being passed correctly.
	games, err := api.GetPlayerMatchHistory("42A9C67482E71FEA", &MatchHistoryOptions{
		Limit:        3,
		Offset:       0,
		CountResults: false,
	})
	if err != nil {
		t.Error("error during `GetPlayerMatchHistory`")
	}
	// The default is 10, this is just for perfomance reasons
	if len(games) != 2 {
		t.Error("error `len(*games)` != 2")
	}
}

func TestGetPlayerMatchHistoryOnErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/matchHistory/42A9C67482E71FEA",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetPlayerMatchHistory("42A9C67482E71FEA", &MatchHistoryOptions{})
	if err == nil {
		t.Error("error GetPlayerMatchHistory``doesn't return error")
	}
}
