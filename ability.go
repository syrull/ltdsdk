package ltdsdk

type Ability struct {
	ID            string `json:"_id"`
	AoeRange      string `json:"aoeRange"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	Duration      string `json:"duration"`
	DescriptionID string `json:"descriptionId"`
	IconPath      string `json:"iconPath"`
	Name          string `json:"name"`
	Tooltip       string `json:"tooltip"`
	Bounces       string `json:"bounces"`
	BaseDamage    string `json:"baseDamage"`
	Cooldown      string `json:"cooldown"`
}
