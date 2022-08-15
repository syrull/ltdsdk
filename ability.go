package ltdsdk

import (
	"fmt"
	"strconv"
)

type abilityResponse struct {
	Id            string `json:"_id"`
	AoeRange      string `json:"aoeRange"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	Duration      string `json:"duration"`
	DescriptionId string `json:"descriptionId"`
	IconPath      string `json:"iconPath"`
	Name          string `json:"name"`
	Tooltip       string `json:"tooltip"`
	Bounces       string `json:"bounces"`
	BaseDamage    string `json:"baseDamage"`
	Cooldown      string `json:"cooldown"`
}

type Ability struct {
	Id          string
	AoeRange    int
	Description string
	Duration    float32
	IconPath    string
	Name        string
	Tooltip     string
	Bounces     int
	BaseDamage  float32
	Cooldown    float32
	Raw         abilityResponse
}

func newAbility(ar *abilityResponse) *Ability {
	aoeRange, err := strconv.Atoi(ar.AoeRange)
	if err != nil {
		aoeRange = 0
	}
	duration, err := strconv.ParseFloat(ar.Duration, 32)
	if err != nil {
		duration = 0
	}
	bounces, err := strconv.Atoi(ar.Bounces)
	if err != nil {
		bounces = 0
	}
	baseDmg, err := strconv.ParseFloat(ar.BaseDamage, 32)
	if err != nil {
		baseDmg = 0
	}
	cd, err := strconv.ParseFloat(ar.Cooldown, 32)
	if err != nil {
		cd = 0
	}
	return &Ability{
		Id:          ar.Id,
		AoeRange:    aoeRange,
		Description: ar.Description,
		Duration:    float32(duration),
		IconPath:    ar.IconPath,
		Name:        ar.Name,
		Tooltip:     ar.Tooltip,
		Bounces:     bounces,
		BaseDamage:  float32(baseDmg),
		Cooldown:    float32(cd),
		Raw:         *ar,
	}
}

func (l *LegionTDSdk) GetAbility(abilityName string) (*Ability, error) {
	abilityResponse := new(abilityResponse)
	endpoint := fmt.Sprintf("info/abilities/byId/%s", abilityName)
	err := l.GetRequest(endpoint, nil, abilityResponse)
	if err != nil {
		return nil, err
	}
	ability := newAbility(abilityResponse)
	return ability, nil
}
