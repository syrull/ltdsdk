package ltdsdk

import (
	"fmt"
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
	return &Ability{
		Id:          ar.Id,
		AoeRange:    parseStringToInt(ar.AoeRange, 0),
		Description: ar.Description,
		Duration:    parseStringToFloat32(ar.Duration, 0),
		IconPath:    ar.IconPath,
		Name:        ar.Name,
		Tooltip:     ar.Tooltip,
		Bounces:     parseStringToInt(ar.Bounces, 0),
		BaseDamage:  parseStringToFloat32(ar.BaseDamage, 0),
		Cooldown:    parseStringToFloat32(ar.Cooldown, 0),
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
