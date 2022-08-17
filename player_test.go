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
