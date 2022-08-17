package ltdsdk

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetUnit(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/units/byName/Atom",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/units/byName_Atom.json")
			return httpmock.NewJsonResponse(200, data)
		})
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/generator_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/abilities/byId_generator_ability_id.json")
			return httpmock.NewJsonResponse(200, data)
		})
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/ionic_force_aspd_ability_id",
		func(_ *http.Request) (*http.Response, error) {
			data := LoadFixture("test_responses/info/abilities/byId_ionic_force_aspd_ability_id.json")
			return httpmock.NewJsonResponse(200, data)
		})

	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	unit, err := api.GetUnit("Atom")
	if err != nil {
		t.Error("error during `GetUnit`")
	}
	if unit.Name != "Atom" {
		t.Error("error `unit.Name` is not `Atom`")
	}
	if unit.ArmorType != "Immaterial" {
		t.Error("error `unit.ArmorType` is not `Immaterial`")
	}
	if unit.Abilities[0].Name != "Generator" {
		t.Error("error `unit.Abilities[0].Name` is not `Generator`")
	}
	if unit.ExpectedDamage != 38.00 {
		t.Error("error `unit.ExpectedDamage` is not `38.00`")
	}
}

func TestGetUnitErrorResponse(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/units/byName/Atom",
		func(_ *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(404, "error"), nil
		})

	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	_, err := api.GetUnit("Atom")
	if err == nil {
		t.Error("error `GetUnit` doesn't return erorr")
	}
}
