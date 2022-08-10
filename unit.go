package ltdsdk

import "log"

type Unit struct {
	SDK                 *LegionTDSdk
	ID                  string   `json:"_id"`
	UnitID              string   `json:"unitId"`
	Version             string   `json:"version"`
	AbilitiesID         []string `json:"abilities"`
	ArmorType           string   `json:"armorType"`
	AspdInverted        string   `json:"aspdInverted"`
	AttackMode          string   `json:"attackMode"`
	AttackRange         string   `json:"attackRange"`
	AttackSpeed         string   `json:"attackSpeed"`
	AttackSpeedInverted string   `json:"attackSpeedInverted"`
	AttackType          string   `json:"attackType"`
	AvgAspd             string   `json:"avgAspd"`     // This is outdated in the Schema
	AvgAspdDiff         string   `json:"avgAspdDiff"` // This is outdated in the Schema
	AvgCost             string   `json:"avgCost"`     // This is outdated in the Schema
	AvgCostDiff         string   `json:"avgCostDiff"` // This is outdated in the Schema
	AvgDmg              string   `json:"avgDmg"`      // This is outdated in the Schema
	AvgDmgDiff          string   `json:"avgDmgDiff"`  // This is outdated in the Schema
	AvgHp               string   `json:"avgHp"`       // This is outdated in the Schema
	AvgHpDiff           string   `json:"avgHpDiff"`   // This is outdated in the Schema
	AvgMspd             string   `json:"avgMspd"`     // This is outdated in the Schema
	AvgMspdDiff         string   `json:"avgMspdDiff"` // This is outdated in the Schema
	CategoryClass       string   `json:"categoryClass"`
	Description         string   `json:"description"`
	DescriptionID       string   `json:"descriptionId"`
	DmgBase             string   `json:"dmgBase"`     // This is outdated in the Schema
	DmgExpected         string   `json:"dmgExpected"` // This is outdated in the Schema
	DmgMax              string   `json:"dmgMax"`      // This is outdated in the Schema
	Dps                 string   `json:"dps"`
	Flags               string   `json:"flags"`
	GoldBounty          string   `json:"goldBounty"`
	GoldCost            string   `json:"goldCost"`
	GoldValue           string   `json:"goldValue"`
	Hp                  string   `json:"hp"`
	IconPath            string   `json:"iconPath"`
	IncomeBonus         string   `json:"incomeBonus"`
	InfoSketchfab       string   `json:"infoSketchfab"`
	InfoTier            string   `json:"infoTier"`
	IsEnabled           bool     `json:"isEnabled"` // This is outdated in the Schema
	LegionID            string   `json:"legionId"`
	ModelScale          string   `json:"modelScale"`
	MoveSpeed           string   `json:"moveSpeed"` // This is outdated in the Schema
	MoveType            string   `json:"moveType"`
	Mp                  string   `json:"mp"`
	MspdText            string   `json:"mspdText"`
	MythiumCost         string   `json:"mythiumCost"`
	Name                string   `json:"name"`
	Radius              string   `json:"radius"`
	RangeText           string   `json:"rangeText"`
	SortOrder           string   `json:"sortOrder"` // This is outdated in the Schema
	SplashPath          string   `json:"splashPath"`
	StockMax            string   `json:"stockMax"`
	StockTime           string   `json:"stockTime"`
	SketchfabUrl        string   `json:"sketchfabUrl"`
	Tooltip             string   `json:"tooltip"`
	TotalValue          string   `json:"totalValue"`
	UnitClass           string   `json:"unitClass"`
	UpgradesFrom        []string `json:"upgradesFrom"`
	UpgradesTo          []string `json:"upgradesTo"`
}

// Getting the abilities using a lazy load for the abilities
func (u *Unit) Abilities() []Ability {
	var abilities = []Ability{}
	for _, a := range u.AbilitiesID {
		ability, err := u.SDK.GetAbility(a)
		if err != nil {
			log.Fatal(err)
		}
		abilities = append(abilities, ability)
	}
	return abilities
}
