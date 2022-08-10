package ltdsdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type LegionTDSdk struct {
	client    *http.Client
	hostUrl   string
	secretKey string
}

func NewLTDSDK(secretKey string, hostUrl string) *LegionTDSdk {
	httpClient := http.DefaultClient
	return &LegionTDSdk{secretKey: secretKey, client: httpClient, hostUrl: hostUrl}
}

func (l *LegionTDSdk) CreateAuthenticatedRequest(method string, url *url.URL) *http.Request {
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-api-key", l.secretKey)
	return req
}

// Get a single unit by Name and version, version can be empty that will
// pull the latest version.
// Official Documentation: https://swagger.legiontd2.com/#/Units/getUnitByName
func (l *LegionTDSdk) GetUnit(unitName string, version string) (Unit, error) {
	unit := &Unit{SDK: l}
	path := "units/byName/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, unitName))
	if err != nil {
		log.Fatal(err)
	}
	if version != "" {
		values := url.Query()
		values.Add("version", version)
		url.RawQuery = values.Encode()
	}
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *unit, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *unit, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&unit); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *unit, nil
}

// Fetching a single ability by id
// Official Documentation: https://swagger.legiontd2.com/#/Info/getAbilitiesById
func (l *LegionTDSdk) GetAbility(abilityName string) (Ability, error) {
	ability := new(Ability)
	path := "info/abilities/byId/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, abilityName))
	if err != nil {
		log.Fatal(err)
	}
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *ability, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *ability, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&ability); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *ability, nil
}

// Fetching a single Spell by id
// Official Documentation: https://swagger.legiontd2.com/#/Info/getSpellsById
func (l *LegionTDSdk) GetSpell(spellId string) (Spell, error) {
	spell := new(Spell)
	path := "info/spells/byId/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, spellId))
	if err != nil {
		log.Fatal(err)
	}
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *spell, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *spell, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&spell); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *spell, nil
}

// Fetching a player by name
// Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerByName
func (l *LegionTDSdk) GetPlayerByName(playerName string) (Player, error) {
	player := new(Player)
	path := "players/byName/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerName))
	if err != nil {
		log.Fatal(err)
	}
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *player, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *player, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *player, nil
}

// Fetching a player by Id
// Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerById
func (l *LegionTDSdk) GetPlayerById(playerId string) (Player, error) {
	player := new(Player)
	path := "players/byId/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerId))
	if err != nil {
		log.Fatal(err)
	}
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *player, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *player, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *player, nil
}

// Fetching a player by Id
// Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerFriends
func (l *LegionTDSdk) GetMostPlayedWith(playerId string, limit int, offset int) (PlayerGames, error) {
	playerGames := new(PlayerGames)
	path := "players/bestFriends/"
	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerId))
	if err != nil {
		log.Fatal(err)
	}
	values := url.Query()
	values.Add("limit", strconv.Itoa(limit))
	values.Add("offset", strconv.Itoa(offset))
	url.RawQuery = values.Encode()
	request := l.CreateAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return *playerGames, err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return *playerGames, message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&playerGames); err != nil {
		log.Fatal("Error occurred during decoding!")
	}
	return *playerGames, nil
}
