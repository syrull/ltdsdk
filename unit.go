package ltdsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type unitResponse struct {
	Id            string   `json:"_id"`
	UnitId        string   `json:"unitId"`
	Version       string   `json:"version"`
	Abilities     []string `json:"abilities"`
	ArmorType     string   `json:"armorType"`
	AspdInverted  string   `json:"aspdInverted"`
	AttackMode    string   `json:"attackMode"`
	AttackRange   string   `json:"attackRange"`
	AttackSpeed   string   `json:"attackSpeed"`
	AttackType    string   `json:"attackType"`
	AvgAspd       string   `json:"avgAspd"`
	AvgAspdDiff   string   `json:"avgAspdDiff"`
	AvgCost       string   `json:"avgCost"`
	AvgCostDiff   string   `json:"avgCostDiff"`
	AvgDmg        string   `json:"avgDmg"`
	AvgDmgDiff    string   `json:"avgDmgDiff"`
	AvgHp         string   `json:"avgHp"`
	AvgHpDiff     string   `json:"avgHpDiff"`
	AvgMspd       string   `json:"avgMspd"`
	AvgMspdDiff   string   `json:"avgMspdDiff"`
	CategoryClass string   `json:"categoryClass"`
	Description   string   `json:"description"`
	DescriptionID string   `json:"descriptionId"`
	DmgBase       string   `json:"dmgBase"`
	DmgExpected   string   `json:"dmgExpected"`
	DmgMax        string   `json:"dmgMax"`
	Dps           string   `json:"dps"`
	Flags         string   `json:"flags"`
	GoldBounty    string   `json:"goldBounty"`
	GoldCost      string   `json:"goldCost"`
	GoldValue     string   `json:"goldValue"`
	Hp            string   `json:"hp"`
	IconPath      string   `json:"iconPath"`
	IncomeBonus   string   `json:"incomeBonus"`
	InfoSketchfab string   `json:"infoSketchfab"`
	InfoTier      string   `json:"infoTier"`
	IsEnabled     bool     `json:"isEnabled"`
	LegionId      string   `json:"legionId"`
	ModelScale    string   `json:"modelScale"`
	MoveSpeed     string   `json:"moveSpeed"`
	MoveType      string   `json:"moveType"`
	Mp            string   `json:"mp"`
	MspdText      string   `json:"mspdText"`
	Name          string   `json:"name"`
	Radius        string   `json:"radius"`
	RangeText     string   `json:"rangeText"`
	SortOrder     string   `json:"sortOrder"`
	SplashPath    string   `json:"splashPath"`
	StockMax      string   `json:"stockMax"`
	StockTime     string   `json:"stockTime"`
	Tooltip       string   `json:"tooltip"`
	TotalValue    string   `json:"totalValue"`
	UnitClass     string   `json:"unitClass"`
	UpgradesFrom  []string `json:"upgradesFrom"`
}

type Unit struct {
	Id                  string
	Version             string
	Abilities           []Ability
	ArmorType           string
	AttackSpeedInverted float32
	AttackMode          string
	AttackRange         int
	AttackType          string
	CategoryClass       string
	Description         string
	BaseDamage          float32
	ExpectedDamage      float32
	MaxDamage           float32
	DamagePerSecond     float32
	Flags               string
	GoldBounty          int
	GoldCost            int
	GoldValue           int
	Health              int
	IsEnabled           bool
	InfoTier            string
	LegionId            string
	ModelScale          float32
	MoveSpeed           float32
	MoveType            string
	Mana                int
	Name                string
	TotalValue          int
	UnitClass           string
	UpgradesFrom        []string
	Raw                 *unitResponse
}

// Creates a new Unit from a response object
func newUnit(ur *unitResponse, l *LegionTDSdk) (*Unit, error) {
	var abilities []Ability
	for _, a := range ur.Abilities {
		ability, err := l.GetAbility(a)
		if err != nil {
			return nil, err
		}
		abilities = append(abilities, *ability)
	}
	return &Unit{
		Id:                  ur.UnitId,
		Version:             ur.Version,
		Abilities:           abilities,
		ArmorType:           ur.ArmorType,
		AttackSpeedInverted: parseStringToFloat32(ur.AspdInverted, 0),
		AttackMode:          ur.AttackMode,
		AttackRange:         parseStringToInt(ur.AttackRange, 0),
		AttackType:          ur.AttackType,
		CategoryClass:       ur.CategoryClass,
		Description:         ur.Description,
		BaseDamage:          parseStringToFloat32(ur.DmgBase, 0),
		ExpectedDamage:      parseStringToFloat32(ur.DmgExpected, 0),
		MaxDamage:           parseStringToFloat32(ur.DmgMax, 0),
		DamagePerSecond:     parseStringToFloat32(ur.Dps, 0),
		Flags:               ur.Flags,
		GoldBounty:          parseStringToInt(ur.GoldBounty, 0),
		GoldCost:            parseStringToInt(ur.GoldCost, 0),
		GoldValue:           parseStringToInt(ur.GoldValue, 0),
		Health:              parseStringToInt(ur.Hp, 0),
		IsEnabled:           ur.IsEnabled,
		InfoTier:            ur.InfoTier,
		LegionId:            ur.LegionId,
		ModelScale:          parseStringToFloat32(ur.ModelScale, 0),
		MoveSpeed:           parseStringToFloat32(ur.MoveSpeed, 0),
		MoveType:            ur.MoveType,
		Mana:                parseStringToInt(ur.Mp, 0),
		Name:                ur.Name,
		TotalValue:          parseStringToInt(ur.TotalValue, 0),
		UnitClass:           ur.UnitClass,
		UpgradesFrom:        ur.UpgradesFrom,
	}, nil
}

// Getting a unit by name, it returns an error in a case where a unit id
// is not being found.
func (l *LegionTDSdk) GetUnit(unitName string) (*Unit, error) {
	unitResp := new(unitResponse)
	endpoint := fmt.Sprintf("units/byName/%s", unitName)
	err := l.getRequest(endpoint, nil, unitResp)
	if err != nil {
		return nil, err
	}
	unit, err := newUnit(unitResp, l)
	if err != nil {
		return nil, err
	}
	return unit, nil
}

// Exporting an Unit to JSON
func (u *Unit) ExportToJson(folder string) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(folder+u.Name+".json", b, 0644)
	if err != nil {
		return err
	}
	return nil
}
