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
	PlayerId                    string      `json:"playerId" db:"id,key,auto"`
	PlayerName                  string      `json:"playerName" db:"player_name"`
	PlayerSlot                  int         `json:"playerSlot" db:"player_slot"`
	Legion                      string      `json:"legion" db:"legion"`
	Workers                     int         `json:"workers" db:"workers"`
	Value                       int         `json:"value" db:"value"`
	Cross                       bool        `json:"cross" db:"cross"`
	GameResult                  string      `json:"gameResult" db:"game_result"`
	OverallElo                  int         `json:"overallElo" db:"overall_elo"`
	ClassicElo                  int         `json:"classicElo" db:"classic_elo"`
	Fighters                    string      `json:"fighters" db:"fighters"`
	Mercenaries                 string      `json:"mercenaries" db:"mercenaries"`
	StayedUntilEnd              bool        `json:"stayedUntilEnd" db:"stayed_until_end"`
	ChosenSpell                 string      `json:"chosenSpell" db:"chosen_spell"`
	ChosenSpellLocation         string      `json:"chosenSpellLocation" db:"chosen_spell_location"`
	PartySize                   int         `json:"partySize" db:"party_size"`
	FirstWaveFighters           string      `json:"firstWaveFighters" db:"first_wave_fighters"`
	Rolls                       string      `json:"rolls" db:"rolls"`
	LegionSpecificElo           int         `json:"legionSpecificElo" db:"legion_specific_elo"`
	PartyMembers                []string    `json:"partyMembers" db:"party_members"`
	PartyMembersIds             []string    `json:"partyMembersIds" db:"party_members_ids"`
	MvpScore                    int         `json:"mvpScore" db:"mvp_score"`
	NetWorthPerWave             []int       `json:"netWorthPerWave" db:"net_worth_per_wave"`
	WorkersPerWave              []int       `json:"workersPerWave" db:"workers_per_wave"`
	IncomePerWave               []int       `json:"incomePerWave" db:"income_per_wave"`
	MercenariesSentPerWave      [][]string  `json:"mercenariesSentPerWave" db:"mvp_score"`
	MercenariesReceivedPerWave  [][]string  `json:"mercenariesReceivedPerWave" db:"mercenaries_received_per_wave"`
	LeaksPerWave                [][]string  `json:"leaksPerWave" db:"leaks_per_wave"`
	BuildPerWave                [][]string  `json:"buildPerWave" db:"build_per_wave"`
	LeakValue                   int         `json:"leakValue" db:"leak_value"`
	LeaksCaughtValue            int         `json:"leaksCaughtValue" db:"leaks_caught_value"`
	LeftAtSeconds               interface{} `json:"leftAtSeconds" db:"left_at_secs"`
	KingUpgradesPerWave         [][]string  `json:"kingUpgradesPerWave" db:"king_upgrades_per_wave"`
	OpponentKingUpgradesPerWave [][]string  `json:"opponentKingUpgradesPerWave" db:"opponent_king_upgrades_per_wave"`
}

type MatchHistoryOptions struct {
	Limit        int  `qs:"limit"`
	Offset       int  `qs:"offset"`
	CountResults bool `qs:"countResults"`
}

// Get a Player by name, it returns an error in a case where a player
// Id is not being found.
func (l *LegionTDSdk) GetPlayerByName(name string) (*Player, error) {
	player := new(Player)
	endpoint := fmt.Sprintf("players/byName/%s", name)
	err := l.getRequest(endpoint, nil, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// Get a player by Id, returns a Player obj pointer, it returns
// an error in a case where a player Id is not being found.
func (l *LegionTDSdk) GetPlayerById(Id string) (*Player, error) {
	player := new(Player)
	endpoint := fmt.Sprintf("players/byId/%s", Id)
	err := l.getRequest(endpoint, nil, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// Getting a player's match history, it accepts MatchHistoryOptions which are
// consisting of Limit=10, Offset=0 and CountResults=false, use the offset
// for pagination.
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
	err := l.getRequest(endpoint, queryString, &games)
	if err != nil {
		return nil, err
	}
	return games, nil
}
