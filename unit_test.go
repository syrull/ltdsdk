package ltdsdk_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/syrull/ltdsdk"
)

func TestGetUnit(t *testing.T) {
	httpmock.Activate()
	data := LoadFixture("test_responses/units/byName_Atom.json")
	ab1 := LoadFixture("test_responses/info/abilities/byId_generator_ability_id.json")
	ab2 := LoadFixture("test_responses/info/abilities/byId_ionic_force_aspd_ability_id.json")

	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/units/byName/Atom",
		httpmock.NewStringResponder(200, data))
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/generator_ability_id",
		httpmock.NewStringResponder(200, ab1))
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/ionic_force_aspd_ability_id",
		httpmock.NewStringResponder(200, ab2))

	api := ltdsdk.NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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
