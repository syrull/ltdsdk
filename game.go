package ltdsdk

import (
	"fmt"
	"time"
)

type Game struct {
	Id                 string    `json:"_id"`
	Version            string    `json:"version"`
	Date               time.Time `json:"date"`
	QueueType          string    `json:"queueType"`
	EndingWave         int       `json:"endingWave"`
	GameLength         int       `json:"gameLength"`
	GameElo            int       `json:"gameElo"`
	PlayerCount        int       `json:"playerCount"`
	HumanCount         int       `json:"humanCount"`
	SpellChoices       []string  `json:"spellChoices"`
	LeftKingPercentHp  []float64 `json:"leftKingPercentHp"`
	RightKingPercentHp []float64 `json:"rightKingPercentHp"`
	KingSpell          string    `json:"kingSpell"`
	PlayersData        []struct {
		PlayerID                    string          `json:"playerId"`
		PlayerName                  string          `json:"playerName"`
		PlayerSlot                  int             `json:"playerSlot"`
		Legion                      string          `json:"legion"`
		Workers                     int             `json:"workers"`
		Value                       int             `json:"value"`
		Cross                       bool            `json:"cross"`
		GameResult                  string          `json:"gameResult"`
		OverallElo                  int             `json:"overallElo"`
		Fighters                    string          `json:"fighters"`
		Mercenaries                 string          `json:"mercenaries"`
		StayedUntilEnd              bool            `json:"stayedUntilEnd"`
		ChosenSpell                 string          `json:"chosenSpell"`
		PartySize                   int             `json:"partySize"`
		FirstWaveFighters           string          `json:"firstWaveFighters"`
		Rolls                       string          `json:"rolls"`
		LegionSpecificElo           int             `json:"legionSpecificElo"`
		PartyMembers                []string        `json:"partyMembers"`
		PartyMembersIds             []string        `json:"partyMembersIds"`
		MvpScore                    int             `json:"mvpScore"`
		NetWorthPerWave             []int           `json:"netWorthPerWave"`
		WorkersPerWave              []int           `json:"workersPerWave"`
		IncomePerWave               []int           `json:"incomePerWave"`
		MercenariesSentPerWave      [][]interface{} `json:"mercenariesSentPerWave"`
		MercenariesReceivedPerWave  [][]interface{} `json:"mercenariesReceivedPerWave"`
		LeaksPerWave                [][]interface{} `json:"leaksPerWave"`
		BuildPerWave                [][]string      `json:"buildPerWave"`
		LeakValue                   int             `json:"leakValue"`
		LeaksCaughtValue            int             `json:"leaksCaughtValue"`
		LeftAtSeconds               float64         `json:"leftAtSeconds"`
		KingUpgradesPerWave         [][]string      `json:"kingUpgradesPerWave"`
		OpponentKingUpgradesPerWave [][]interface{} `json:"opponentKingUpgradesPerWave"`
	} `json:"playersData,omitempty"`
}

type GameOptions struct {
	Version        string `qs:"version"`
	Limit          int    `qs:"limit"`
	Offset         int    `qs:"offset"`
	SortBy         string `qs:"sortBy"`
	SortDirection  int    `qs:"sortDirection"`
	AfterDate      string `qs:"afterDate"`
	BeforeDate     string `qs:"beforeDate"`
	IncludeDetails bool   `qs:"includeDetails"`
	QueueType      string `qs:"queueType"`
}

func (l *LegionTDSdk) GetGameById(Id string) (*Game, error) {
	game := new(Game)
	endpoint := fmt.Sprintf("games/byId/%s", Id)
	err := l.GetRequest(endpoint, nil, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (l *LegionTDSdk) GetGames(gameOpts *GameOptions) ([]*Game, error) {
	var games []*Game
	gameOptsDef := &GameOptions{}
	// Set the default search options if strcut is empty
	if gameOpts == gameOptsDef {
		gameOpts = &GameOptions{
			Version:        "",
			Limit:          50,
			Offset:         0,
			SortBy:         "date",
			SortDirection:  1,
			AfterDate:      time.Now().Format("2000-01-01 00:00:00"),
			BeforeDate:     time.Now().Format("2000-01-01 00:00:00"),
			IncludeDetails: false,
			QueueType:      "",
		}
	}
	queryString := toQueryString(gameOpts)
	err := l.GetRequest("games/", queryString, &games)
	if err != nil {
		return nil, err
	}
	return games, nil
}
