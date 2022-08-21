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

// Getting a spell by Id, it returns an error in a case where a spell id
// is not being found.
func (l *LegionTDSdk) GetSpell(Id string) (*Spell, error) {
	spell := new(Spell)
	endpoint := fmt.Sprintf("info/spells/byId/%s", Id)
	err := l.getRequest(endpoint, nil, spell)
	if err != nil {
		return nil, err
	}
	return spell, nil
}
