package ltdsdk

import (
	"fmt"
)

type waveResponse struct {
	Id           string `json:"_id"`
	Amount       string `json:"amount"`
	Amount2      string `json:"amount2,omitempty"`
	Category     string `json:"category"`
	IconPath     string `json:"iconPath"`
	LevelNum     string `json:"levelNum"`
	Name         string `json:"name"`
	PrepareTime  string `json:"prepareTime"`
	SpellUnit2Id string `json:"spellUnit2Id,omitempty"`
	TotalReward  string `json:"totalReward"`
	WaveUnitId   string `json:"waveUnitId"`
}

type Wave struct {
	Id           string
	Amount       int
	Amount2      int
	Category     string
	IconPath     string
	LevelNum     int
	Name         string
	PrepareTime  int
	SpellUnit2Id string
	TotalReward  int
	WaveUnitId   string
}

// Creates a new Wave object with defaults from a response
func newWave(wr *waveResponse) *Wave {
	return &Wave{
		Id:           wr.Id,
		Amount:       parseStringToInt(wr.Amount, 0),
		Amount2:      parseStringToInt(wr.Amount2, 0),
		Category:     wr.Category,
		IconPath:     wr.IconPath,
		LevelNum:     parseStringToInt(wr.LevelNum, 0),
		Name:         wr.Name,
		PrepareTime:  parseStringToInt(wr.PrepareTime, 0),
		SpellUnit2Id: wr.SpellUnit2Id,
		TotalReward:  parseStringToInt(wr.TotalReward, 0),
		WaveUnitId:   wr.WaveUnitId,
	}
}

// Getting a Wave by Id
func (l *LegionTDSdk) GetWave(Id string) (*Wave, error) {
	wr := new(waveResponse)
	endpoint := fmt.Sprintf("info/waves/byId/%s", Id)
	err := l.getRequest(endpoint, nil, wr)
	if err != nil {
		return nil, err
	}
	wave := newWave(wr)
	return wave, nil
}
