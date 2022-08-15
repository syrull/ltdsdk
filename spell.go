package ltdsdk

import "fmt"

type Spell struct {
	Id       string `json:"_id"`
	Category string `json:"category"`
	Enabled  bool   `json:"enabled"`
	IconPath string `json:"iconPath"`
	Name     string `json:"name"`
	Tooltip  string `json:"tooltip"`
}

func (l *LegionTDSdk) GetSpell(Id string) (*Spell, error) {
	spell := new(Spell)
	endpoint := fmt.Sprintf("info/spells/byId/%s", Id)
	err := l.GetRequest(endpoint, nil, spell)
	if err != nil {
		return nil, err
	}
	return spell, nil
}
