package ltdsdk

import (
	"fmt"
)

type Description struct {
	Id          string `json:"_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Getting a Description by Id, returns an error if not found.
func (l *LegionTDSdk) GetDescription(Id string) (*Description, error) {
	desc := new(Description)
	endpoint := fmt.Sprintf("info/descriptions/%s", Id)
	err := l.getRequest(endpoint, nil, desc)
	if err != nil {
		return nil, err
	}
	return desc, nil
}
