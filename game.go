package ltdsdk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Game struct {
	Id                 string       `json:"_id" db:"id,key,auto"`
	Version            string       `json:"version" db:"version"`
	Date               time.Time    `json:"date" db:"date"`
	QueueType          string       `json:"queueType" db:"queue_type"`
	EndingWave         int          `json:"endingWave" db:"ending_wave"`
	GameLength         int          `json:"gameLength" db:"game_length"`
	GameElo            int          `json:"gameElo" db:"game_elo"`
	PlayerCount        int          `json:"playerCount" db:"player_count"`
	HumanCount         int          `json:"humanCount" db:"human_count"`
	SpellChoices       []string     `json:"spellChoices" db:"spell_choices"`
	LeftKingPercentHp  []float64    `json:"leftKingPercentHp" db:"left_king_percent_hp"`
	RightKingPercentHp []float64    `json:"rightKingPercentHp" db:"right_king_percent_hp"`
	KingSpell          string       `json:"kingSpell" db:"king_spell"`
	PlayersData        []PlayerData `json:"playersData,omitempty" db:"players_data"`
}

type GameCollection struct {
	Games []*Game
}

type GameOptions struct {
	Version        string `qs:"version"`
	Limit          int    `qs:"limit"`
	Offset         int    `qs:"offset"`
	SortBy         string `qs:"sortBy"`
	SortDirection  int    `qs:"sortDirection"`
	AfterDate      string `qs:"dateAfter"`
	BeforeDate     string `qs:"dateBefore"`
	IncludeDetails bool   `qs:"includeDetails"`
	QueueType      string `qs:"queueType"`
}

// Getting a game by Id, returns an error if not found.
func (l *LegionTDSdk) GetGameById(Id string) (*Game, error) {
	game := new(Game)
	endpoint := fmt.Sprintf("games/byId/%s", Id)
	err := l.getRequest(endpoint, nil, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

// Get latest (default=50) games, the method accpets GameOptions
// which consists of Version, Limit, Offset, SortBy, SortDirection,
// AfterDate, BeforeDate, IncludedDetails, QueueType
// For more information about the GameOptions:
// https://swagger.legiontd2.com/#/Games/getMatchesByFilter
func (l *LegionTDSdk) GetGames(gameOpts *GameOptions) (*GameCollection, error) {
	var games GameCollection
	var gameOptsDef GameOptions
	// Set the default search options if strcut is empty
	if *gameOpts == gameOptsDef {
		gameOpts = &GameOptions{
			Version:        "",
			Limit:          50,
			Offset:         0,
			SortBy:         "date",
			SortDirection:  -1,
			IncludeDetails: false,
			QueueType:      "",
		}
	}
	queryString := toQueryString(gameOpts)
	err := l.getRequest("games", queryString, &games.Games)
	if err != nil {
		return nil, err
	}
	return &games, nil
}

// Exporting a GameCollection into a JSON file for further statistic
func (gc *GameCollection) ExportToJson(outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	err = enc.Encode(gc.Games)
	if err != nil {
		return err
	}
	return nil
}
