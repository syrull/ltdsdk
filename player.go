package ltdsdk

type Player struct {
	ID          string `json:"_id"`
	AvatarURL   string `json:"avatarUrl"`
	Name        string `json:"playerName"`
	GuildAvatar string `json:"guildAvatar"`
	GuildName   string `json:"guildName"`
	GuildTag    string `json:"guildTag"`
}

type PlayerGames []struct {
	ID struct {
		PlayerID   string `json:"playerId"`
		PlayerName string `json:"playerName"`
	} `json:"_id"`
	Count int `json:"count"`
}
