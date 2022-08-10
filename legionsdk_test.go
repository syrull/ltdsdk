package ltdsdk

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetUnit(t *testing.T) {
	httpmock.Activate()
	data := `{
	  "_id": "62dcaf205e725edeeb031685",
	  "unitId": "pollywog_unit_id",
	  "version": "9.06.4.hf1",
	  "abilities": [],
	  "armorType": "Swift",
	  "aspdInverted": "1.05",
	  "attackMode": "Ranged",
	  "attackRange": "350",
	  "attackSpeed": "0.9500",
	  "attackType": "Magic",
	  "avgAspd": "1.01",
	  "avgAspdDiff": "3.51",
	  "avgCost": "18.33",
	  "avgCostDiff": "-18.18",
	  "avgDmg": "6.56",
	  "avgDmgDiff": "6.78",
	  "avgHp": "135.56",
	  "avgHpDiff": "-40.98",
	  "avgMspd": "274.92",
	  "avgMspdDiff": "9.12",
	  "categoryClass": "Standard",
	  "description": "Flying. Can upgrade into two different forms.",
	  "descriptionId": "pollywog_unit_description",
	  "dmgBase": "7",
	  "dmgExpected": "7.00",
	  "dmgMax": "7",
	  "dps": "7.37",
	  "flags": "flags_flying",
	  "goldBounty": "",
	  "goldCost": "15",
	  "goldValue": "",
	  "hp": "80",
	  "iconPath": "Icons/Pollywog.png",
	  "incomeBonus": "",
	  "infoSketchfab": "https://sketchfab.com/models/baa8b551258947bca5db5ea99758f676",
	  "infoTier": "Tier-1",
	  "isEnabled": true,
	  "legionId": "atlantean_legion_id",
	  "modelScale": "1.0000",
	  "moveSpeed": "300",
	  "moveType": "Air",
	  "mp": "",
	  "mspdText": "Average",
	  "name": "Pollywog",
	  "radius": "Tiny",
	  "rangeText": "Medium",
	  "sortOrder": "atlantean_legion_id.T1.15.Pollywog",
	  "splashPath": "Splashes/Pollywog.png",
	  "stockMax": "",
	  "stockTime": "",
	  "tooltip": "Flying. Can upgrade into two different forms.",
	  "totalValue": "15",
	  "unitClass": "Fighter",
	  "upgradesFrom": []
	}`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/units/byName/Pollywog",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	unit, err := api.GetUnit("Pollywog", "9.06.4")
	if err != nil {
		t.Error("error during `GetUnit`")
	}
	if unit.Name != "Pollywog" {
		t.Error("error `unit.Name` is not `Pollywog`")
	}
	if unit.ArmorType != "Swift" {
		t.Error("error `unit.ArmorType` is not `Swift`")
	}
}

func TestGetAbility(t *testing.T) {
	httpmock.Activate()
	data := `{
		"_id": "pulverize_melee_ability_id",
		"bounces": "0",
		"category": "abilities",
		"description": "This unit is a Pulverizer. It has a $1% ($2% for ranged) chance to stun its target for $3 seconds ($4 to bosses)",
		"descriptionId": "pulverize_description",
		"duration": "2.00000",
		"iconPath": "Icons/Pulverize.png",
		"name": "Pulverize",
		"tooltip": "This unit is a Pulverizer. It has a 34% (17% for ranged) chance to stun its target for 2 seconds (0.3 to bosses)"
	}`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/abilities/byId/pulverize_melee_ability_id",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	ability, err := api.GetAbility("pulverize_melee_ability_id")
	if err != nil {
		t.Error("error during `GetAbility`")
	}
	if ability.Name != "Pulverize" {
		t.Error("error `ability.Name` is not `Pulverize`")
	}
	if ability.Duration != "2.00000" {
		t.Error("error `ability.Duration` is not `2.00000`")
	}
}

func TestGetSpell(t *testing.T) {
	httpmock.Activate()
	data := `{
		"_id": "allowance_powerup_id",
		"category": "spells",
		"enabled": true,
		"iconPath": "Icons/Allowance.png",
		"name": "Allowance",
		"tooltip": "+100 gold"
	  }`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/info/spells/byId/allowance_powerup_id",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	spell, err := api.GetSpell("allowance_powerup_id")
	if err != nil {
		t.Error("error during `GetSpell`")
	}
	if spell.Name != "Allowance" {
		t.Error("error `spell.Name` is not `Allowance`")
	}
	if spell.Tooltip != "+100 gold" {
		t.Error("error `spell.Tooltip` is not `+100 gold`")
	}
}

func TestGetPlayerByName(t *testing.T) {
	httpmock.Activate()
	data := `{
		"_id": "42A9C67482E71FEA",
		"avatarUrl": "Icons/Atom.png",
		"playerName": "syll",
		"guildAvatar": "",
		"guildName": "",
		"guildTag": ""
	  }`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byName/syll",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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
	data := `{
		"_id": "42A9C67482E71FEA",
		"avatarUrl": "Icons/Atom.png",
		"playerName": "syll",
		"guildAvatar": "",
		"guildName": "",
		"guildTag": ""
	  }`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/byId/42A9C67482E71FEA",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
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

func TestGetMostPlayedWith(t *testing.T) {
	httpmock.Activate()
	data := `[
		{
		  "_id": {
			"playerId": "42A9C67482E71FEA",
			"playerName": "test1"
		  },
		  "count": 3
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FEB",
			"playerName": "test2"
		  },
		  "count": 3
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FEC",
			"playerName": "test3"
		  },
		  "count": 3
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FED",
			"playerName": "test4"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FEE",
			"playerName": "test5"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FEF",
			"playerName": "test6"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FE3",
			"playerName": "test7"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FE1",
			"playerName": "test8"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FE9",
			"playerName": "test9"
		  },
		  "count": 2
		},
		{
		  "_id": {
			"playerId": "42A9C67482E71FEY",
			"playerName": "test10"
		  },
		  "count": 2
		}
	  ]`
	httpmock.RegisterResponder("GET", "https://apiv2.legiontd2.com/players/bestFriends/42A9C67482E71FEA?limit=10&offset=0",
		httpmock.NewStringResponder(200, data))
	api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
	playerGames, err := api.GetMostPlayedWith("42A9C67482E71FEA", 10, 0)
	for game := range playerGames {
		fmt.Println(game)
	}
	if err != nil {
		t.Error("error during `GetPlayerByName`")
	}
	if len(playerGames) != 10 {
		t.Error("error `len(playerGames)` is not `10`")
	}
	if playerGames[0].ID.PlayerName != "test1" {
		t.Error("error `playerGames[0].ID.PlayerName` is not `test1`")
	}
}
