package ltdsdk

import (
	"fmt"
	"time"
)

type Game struct {
	Id                 string       `json:"_id"`
	Version            string       `json:"version"`
	Date               time.Time    `json:"date"`
	QueueType          string       `json:"queueType"`
	EndingWave         int          `json:"endingWave"`
	GameLength         int          `json:"gameLength"`
	GameElo            int          `json:"gameElo"`
	PlayerCount        int          `json:"playerCount"`
	HumanCount         int          `json:"humanCount"`
	SpellChoices       []string     `json:"spellChoices"`
	LeftKingPercentHp  []float64    `json:"leftKingPercentHp"`
	RightKingPercentHp []float64    `json:"rightKingPercentHp"`
	KingSpell          string       `json:"kingSpell"`
	PlayersData        []PlayerData `json:"playersData,omitempty"`
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
	var gameOptsDef GameOptions
	// Set the default search options if strcut is empty
	if *gameOpts == gameOptsDef {
		gameOpts = &GameOptions{
			Version:        "",
			Limit:          50,
			Offset:         0,
			SortBy:         "date",
			SortDirection:  1,
			AfterDate:      time.Now().Add(-24 * time.Hour).Format("2000-01-01 00:00:00"),
			BeforeDate:     time.Now().Format("2000-01-01 00:00:00"),
			IncludeDetails: false,
			QueueType:      "",
		}
	}
	queryString := toQueryString(gameOpts)
	err := l.GetRequest("games", queryString, &games)
	if err != nil {
		return nil, err
	}
	return games, nil
}
