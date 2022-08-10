package ltdsdk

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestUnitGetAbilities(t *testing.T) {
	httpmock.Activate()
	dataUnit := `{
		"_id": "62dcaf205e725edeeb031743",
		"unitId": "oathbreaker_unit_id",
		"version": "9.06.4.hf1",
		"abilities": [
		  "unchained_rage_activation_ability_id"
		],
		"armorType": "Fortified",
		"aspdInverted": ".93",
		"attackMode": "Melee",
		"attackRange": "100",
		"attackSpeed": "1.08000",
		"attackType": "Impact",
		"avgAspd": "1.01",
		"avgAspdDiff": "-8.32",
		"avgCost": "119.06",
		"avgCostDiff": "-20.21",
		"avgDmg": "40.44",
		"avgDmgDiff": "-33.23",
		"avgHp": "927.50",
		"avgHpDiff": "21.83",
		"avgMspd": "274.92",
		"avgMspdDiff": "9.12",
		"categoryClass": "Standard",
		"description": "Tanky. Attacks quickly when low life.",
		"descriptionId": "oathbreaker_unit_description",
		"dmgBase": "27",
		"dmgExpected": "27.00",
		"dmgMax": "27",
		"dps": "25.00",
		"flags": "flags_ground,flags_organic",
		"goldBounty": "",
		"goldCost": "75",
		"goldValue": "",
		"hp": "1130",
		"iconPath": "Icons/Oathbreaker.png",
		"incomeBonus": "",
		"infoSketchfab": "https://sketchfab.com/models/3cd702c974dd485791db52f5894a8d91",
		"infoTier": "Tier-1",
		"isEnabled": true,
		"legionId": "divine_legion_id",
		"modelScale": "1.15000",
		"moveSpeed": "300",
		"moveType": "Ground",
		"mp": "",
		"mspdText": "Average",
		"name": "Oathbreaker",
		"radius": "Tiny",
		"rangeText": "Melee",
		"sortOrder": "divine_legion_id.T1U.75.Oathbreaker",
		"splashPath": "Splashes/Oathbreaker.png",
		"stockMax": "",
		"stockTime": "",
		"tooltip": "Tanky. Attacks quickly when low life.",
		"totalValue": "95",
		"unitClass": "Fighter",
		"upgradesFrom": [
		  "units chained_fist_unit_id"
		]
	  }`
	dataAbility := `{
		"_id": "unchained_rage_activation_ability_id",
		"category": "abilities",
		"description": "Once it reaches $1% life, it attacks $2% faster, but hits random targets.",
		"descriptionId": "unchained_rage_description",
		"iconPath": "Icons/UnchainedRage.png",
		"name": "Unchained Rage",
		"tooltip": "Once it reaches 50% life, it attacks 275% faster, but hits random targets."
	  }`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/units/byName/Oathbreaker",
		httpmock.NewStringResponder(200, dataUnit))
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/unchained_rage_activation_ability_id",
		httpmock.NewStringResponder(200, dataAbility))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	unit, err := api.GetUnit("Oathbreaker", "")
	if err != nil {
		t.Error("error during `GetUnit`")
	}
	if unit.Name != "Oathbreaker" {
		t.Error("error `unit.Name` is not `Oathbreaker`")
	}
	if unit.ArmorType != "Fortified" {
		t.Error("error `unit.ArmorType` is not `Fortified`")
	}
	if abilities := unit.Abilities(); len(abilities) == 1 {
		if abilities[0].Name != "Unchained Rage" {
			t.Error("error `abilities[0].Name` is not `Unchained Rage`")
		}
	} else {
		t.Error("error `unit.Abilities()` is not `1`")
	}
}
