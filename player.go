package ltdsdk

import "fmt"

type Player struct {
	ID          string `json:"_id"`
	AvatarURL   string `json:"avatarUrl"`
	Name        string `json:"playerName"`
	GuildAvatar string `json:"guildAvatar"`
	GuildName   string `json:"guildName"`
	GuildTag    string `json:"guildTag"`
}

func (l *LegionTDSdk) GetPlayerByName(name string) (*Player, error) {
	player := new(Player)
	endpoint := fmt.Sprintf("players/byName/%s", name)
	err := l.GetRequest(endpoint, nil, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (l *LegionTDSdk) GetPlayerById(Id string) (*Player, error) {
	player := new(Player)
	endpoint := fmt.Sprintf("players/byId/%s", Id)
	err := l.GetRequest(endpoint, nil, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}
