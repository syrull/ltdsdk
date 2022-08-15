package ltdsdk

import (
	"fmt"
	"strconv"
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

func newUnit(ur *unitResponse, l *LegionTDSdk) (*Unit, error) {
	var abilities []Ability
	unit := new(Unit)
	for _, a := range ur.Abilities {
		ability, err := l.GetAbility(a)
		if err != nil {
			return unit, err
		}
		abilities = append(abilities, *ability)
	}
	attackSpeedInverted, err := strconv.ParseFloat(ur.AspdInverted, 32)
	if err != nil {
		attackSpeedInverted = 0
	}
	attackRange, err := strconv.Atoi(ur.AttackRange)
	if err != nil {
		attackRange = 0
	}
	baseDmg, err := strconv.ParseFloat(ur.DmgBase, 32)
	if err != nil {
		baseDmg = 0
	}
	expDmg, err := strconv.ParseFloat(ur.DmgExpected, 32)
	if err != nil {
		expDmg = 0
	}
	maxDmg, err := strconv.ParseFloat(ur.DmgMax, 32)
	if err != nil {
		maxDmg = 0
	}
	dps, err := strconv.ParseFloat(ur.Dps, 32)
	if err != nil {
		dps = 0
	}
	goldBounty, err := strconv.Atoi(ur.GoldBounty)
	if err != nil {
		goldBounty = 0
	}
	goldCost, err := strconv.Atoi(ur.GoldCost)
	if err != nil {
		goldCost = 0
	}
	goldValue, err := strconv.Atoi(ur.GoldValue)
	if err != nil {
		goldValue = 0
	}
	health, err := strconv.Atoi(ur.Hp)
	if err != nil {
		health = 0
	}
	modelScale, err := strconv.ParseFloat(ur.ModelScale, 32)
	if err != nil {
		modelScale = 0
	}
	moveSpeed, err := strconv.ParseFloat(ur.MoveSpeed, 32)
	if err != nil {
		moveSpeed = 0
	}
	mana, err := strconv.Atoi(ur.Mp)
	if err != nil {
		mana = 0
	}
	totalValue, err := strconv.Atoi(ur.TotalValue)
	if err != nil {
		totalValue = 0
	}
	return &Unit{
		Id:                  ur.UnitId,
		Version:             ur.Version,
		Abilities:           abilities,
		ArmorType:           ur.ArmorType,
		AttackSpeedInverted: float32(attackSpeedInverted),
		AttackMode:          ur.AttackMode,
		AttackRange:         attackRange,
		AttackType:          ur.AttackType,
		CategoryClass:       ur.CategoryClass,
		Description:         ur.Description,
		BaseDamage:          float32(baseDmg),
		ExpectedDamage:      float32(expDmg),
		MaxDamage:           float32(maxDmg),
		DamagePerSecond:     float32(dps),
		Flags:               ur.Flags,
		GoldBounty:          goldBounty,
		GoldCost:            goldCost,
		GoldValue:           goldValue,
		Health:              health,
		IsEnabled:           ur.IsEnabled,
		InfoTier:            ur.InfoTier,
		LegionId:            ur.LegionId,
		ModelScale:          float32(modelScale),
		MoveSpeed:           float32(moveSpeed),
		MoveType:            ur.MoveType,
		Mana:                mana,
		Name:                ur.Name,
		TotalValue:          totalValue,
		UnitClass:           ur.UnitClass,
		UpgradesFrom:        ur.UpgradesFrom,
	}, nil
}

func (l *LegionTDSdk) GetUnit(unitName string) (*Unit, error) {
	unitResp := new(unitResponse)
	endpoint := fmt.Sprintf("units/byName/%s", unitName)
	err := l.GetRequest(endpoint, nil, unitResp)
	if err != nil {
		return nil, err
	}
	unit, err := newUnit(unitResp, l)
	if err != nil {
		return nil, err
	}
	return unit, nil
}
