package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetWave(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/waves/byId/level21_wave_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/waves/byId_level21_wave_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	wave, err := api.GetWave("level21_wave_id")
	if err != nil {
		t.Error("error during `GetAbility`")
	}
	if wave.Category != "waves" {
		t.Error("error `wave.Category` is not `waves`")
	}
	if wave.Amount != 10 {
		t.Error("error `wave.Amount` is not 10")
	}
	if wave.Amount2 != 1 {
		t.Error("error `wave.Amount2` is not 1")
	}
}

func TestGetWaveErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/waves/byId/level21_wave_id",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetWave("level21_wave_id")
	if err == nil {
		t.Error("error `GetWave` doesn't return error")
	}
}
