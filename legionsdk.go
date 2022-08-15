package ltdsdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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

func (l *LegionTDSdk) createAuthenticatedRequest(method string, url *url.URL) *http.Request {
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-api-key", l.secretKey)
	return req
}

func (l *LegionTDSdk) GetRequest(endpoint string, queryString map[string]string, obj interface{}) error {
	url, err := url.Parse(fmt.Sprintf("%s%s", l.hostUrl, endpoint))
	if err != nil {
		return err
	}
	if queryString != nil {
		values := url.Query()
		for k, v := range queryString {
			values.Add(k, v)
		}
	}
	request := l.createAuthenticatedRequest("GET", url)
	resp, err := l.client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		message := fmt.Errorf("API returned %v", resp.StatusCode)
		return message
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return err
	}
	return nil
}

// // Fetching a single Spell by id
// // Official Documentation: https://swagger.legiontd2.com/#/Info/getSpellsById
// func (l *LegionTDSdk) GetSpell(spellId string) (ltdsdk.SpellResponse, error) {
// 	spell := new(Spell)
// 	path := "info/spells/byId/"
// 	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, spellId))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	request := l.createAuthenticatedRequest("GET", url)
// 	resp, err := l.client.Do(request)
// 	if err != nil {
// 		return *spell, err
// 	}
// 	if resp.StatusCode != 200 {
// 		message := fmt.Errorf("API returned %v", resp.StatusCode)
// 		return *spell, message
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}(resp.Body)
// 	if err := json.NewDecoder(resp.Body).Decode(&spell); err != nil {
// 		log.Fatal("Error occurred during decoding!")
// 	}
// 	return *spell, nil
// // }

// // Fetching a player by name
// // Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerByName
// func (l *LegionTDSdk) GetPlayerByName(playerName string) (ltdsdk.Player, error) {
// 	player := new(Player)
// 	path := "players/byName/"
// 	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerName))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	request := l.createAuthenticatedRequest("GET", url)
// 	resp, err := l.client.Do(request)
// 	if err != nil {
// 		return *player, err
// 	}
// 	if resp.StatusCode != 200 {
// 		message := fmt.Errorf("API returned %v", resp.StatusCode)
// 		return *player, message
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}(resp.Body)
// 	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
// 		log.Fatal("Error occurred during decoding!")
// 	}
// 	return *player, nil
// }

// // Fetching a player by Id
// // Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerById
// func (l *LegionTDSdk) GetPlayerById(playerId string) (ltdsdk.Player, error) {
// 	player := new(Player)
// 	path := "players/byId/"
// 	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerId))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	request := l.createAuthenticatedRequest("GET", url)
// 	resp, err := l.client.Do(request)
// 	if err != nil {
// 		return *player, err
// 	}
// 	if resp.StatusCode != 200 {
// 		message := fmt.Errorf("API returned %v", resp.StatusCode)
// 		return *player, message
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}(resp.Body)
// 	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
// 		log.Fatal("Error occurred during decoding!")
// 	}
// 	return *player, nil
// }

// Fetching a player by Id
// Official Documentation: https://swagger.legiontd2.com/#/Player/getPlayerFriends
// func (l *LegionTDSdk) GetMostPlayedWith(playerId string, limit int, offset int) (PlayerGames, error) {
// 	playerGames := new(PlayerGames)
// 	path := "players/bestFriends/"
// 	url, err := url.Parse(fmt.Sprintf("%s%s%s", l.hostUrl, path, playerId))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	values := url.Query()
// 	values.Add("limit", strconv.Itoa(limit))
// 	values.Add("offset", strconv.Itoa(offset))
// 	url.RawQuery = values.Encode()
// 	request := l.createAuthenticatedRequest("GET", url)
// 	resp, err := l.client.Do(request)
// 	if err != nil {
// 		return *playerGames, err
// 	}
// 	if resp.StatusCode != 200 {
// 		message := fmt.Errorf("API returned %v", resp.StatusCode)
// 		return *playerGames, message
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}(resp.Body)
// 	if err := json.NewDecoder(resp.Body).Decode(&playerGames); err != nil {
// 		log.Fatal("Error occurred during decoding!")
// 	}
// 	return *playerGames, nil
// }
