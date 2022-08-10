package ltdsdk

type Spell struct {
	ID       string `json:"_id"`
	Category string `json:"category"`
	Enabled  bool   `json:"enabled"`
	IconPath string `json:"iconPath"`
	Name     string `json:"name"`
	Tooltip  string `json:"tooltip"`
}
