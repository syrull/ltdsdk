package ltdsdk

import "fmt"

type Player struct {
	Id          string `json:"_id"`
	AvatarURL   string `json:"avatarUrl"`
	Name        string `json:"playerName"`
	GuildAvatar string `json:"guildAvatar"`
	GuildName   string `json:"guildName"`
	GuildTag    string `json:"guildTag"`
}

type PlayerData struct {
	PlayerId                    string      `json:"playerId"`
	PlayerName                  string      `json:"playerName"`
	PlayerSlot                  int         `json:"playerSlot"`
	Legion                      string      `json:"legion"`
	Workers                     int         `json:"workers"`
	Value                       int         `json:"value"`
	Cross                       bool        `json:"cross"`
	GameResult                  string      `json:"gameResult"`
	OverallElo                  int         `json:"overallElo"`
	ClassicElo                  int         `json:"classicElo"`
	Fighters                    string      `json:"fighters"`
	Mercenaries                 string      `json:"mercenaries"`
	StayedUntilEnd              bool        `json:"stayedUntilEnd"`
	ChosenSpell                 string      `json:"chosenSpell"`
	ChosenSpellLocation         string      `json:"chosenSpellLocation"`
	PartySize                   int         `json:"partySize"`
	FirstWaveFighters           string      `json:"firstWaveFighters"`
	Rolls                       string      `json:"rolls"`
	LegionSpecificElo           int         `json:"legionSpecificElo"`
	PartyMembers                []string    `json:"partyMembers"`
	PartyMembersIds             []string    `json:"partyMembersIds"`
	MvpScore                    int         `json:"mvpScore"`
	NetWorthPerWave             []int       `json:"netWorthPerWave"`
	WorkersPerWave              []int       `json:"workersPerWave"`
	IncomePerWave               []int       `json:"incomePerWave"`
	MercenariesSentPerWave      [][]string  `json:"mercenariesSentPerWave"`
	MercenariesReceivedPerWave  [][]string  `json:"mercenariesReceivedPerWave"`
	LeaksPerWave                [][]string  `json:"leaksPerWave"`
	BuildPerWave                [][]string  `json:"buildPerWave"`
	LeakValue                   int         `json:"leakValue"`
	LeaksCaughtValue            int         `json:"leaksCaughtValue"`
	LeftAtSeconds               interface{} `json:"leftAtSeconds"`
	KingUpgradesPerWave         [][]string  `json:"kingUpgradesPerWave"`
	OpponentKingUpgradesPerWave [][]string  `json:"opponentKingUpgradesPerWave"`
}

type MatchHistoryOptions struct {
	Limit        int  `qs:"limit"`
	Offset       int  `qs:"offset"`
	CountResults bool `qs:"countResults"`
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

func (l *LegionTDSdk) GetPlayerMatchHistory(Id string, matchOpts *MatchHistoryOptions) ([]*Game, error) {
	var games []*Game
	var matchOptsDef MatchHistoryOptions
	// Set the default search options if strcut is empty
	if *matchOpts == matchOptsDef {
		matchOpts = &MatchHistoryOptions{
			Limit:        10,
			Offset:       0,
			CountResults: false,
		}
	}
	queryString := toQueryString(matchOpts)
	endpoint := fmt.Sprintf("players/matchHistory/%s", Id)
	err := l.GetRequest(endpoint, queryString, &games)
	if err != nil {
		return nil, err
	}
	return games, nil
}
