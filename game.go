package ltdsdk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
	"github.com/samonzeweb/godb/tablenamer"
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

// Exporting a GameCollection into a JSON file for further processing
func (gc *GameCollection) ExportToJson(outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.Encode(gc.Games)
	return nil
}

// Exporting a GameCollection into a SQLite for further processing
func (gc *GameCollection) ExportToSql(outputDb string) error {
	db, err := godb.Open(sqlite.Adapter, outputDb)
	db.SetDefaultTableNamer(tablenamer.Plural())
	if err != nil {
		fmt.Println(err)
	}
	for _, game := range gc.Games {
		_, err := db.InsertInto("games").
			Columns(
				"gameId",
				"version",
				"date",
				"queue_type",
				"ending_wave",
				"game_length",
				"game_elo",
				"player_count",
				"human_count",
				"spell_choices_csv",
				"left_king_percent_hp_csv",
				"right_king_percent_hp_csv",
				"king_spell",
			).
			Values(
				game.Id,
				game.Version,
				game.Date,
				game.QueueType,
				game.EndingWave,
				game.GameLength,
				game.GameElo,
				game.PlayerCount,
				game.HumanCount,
				sliceToCsv(game.SpellChoices),
				sliceToCsv(game.LeftKingPercentHp),
				sliceToCsv(game.RightKingPercentHp),
				game.KingSpell,
			).
			Do()
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
