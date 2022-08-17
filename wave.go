package ltdsdk

import (
	"fmt"
	"strconv"
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

func newWave(wr *waveResponse) (*Wave, error) {
	amount, err := strconv.Atoi(wr.Amount)
	if err != nil {
		return nil, err
	}
	amount2, err := strconv.Atoi(wr.Amount2)
	if err != nil {
		return nil, err
	}
	levelNum, err := strconv.Atoi(wr.Amount2)
	if err != nil {
		return nil, err
	}
	prepareTime, err := strconv.Atoi(wr.PrepareTime)
	if err != nil {
		return nil, err
	}
	totalReward, err := strconv.Atoi(wr.TotalReward)
	if err != nil {
		return nil, err
	}
	return &Wave{
		Id:           wr.Id,
		Amount:       amount,
		Amount2:      amount2,
		Category:     wr.Category,
		IconPath:     wr.IconPath,
		LevelNum:     levelNum,
		Name:         wr.Name,
		PrepareTime:  prepareTime,
		SpellUnit2Id: wr.SpellUnit2Id,
		TotalReward:  totalReward,
		WaveUnitId:   wr.WaveUnitId,
	}, nil
}

func (l *LegionTDSdk) GetWave(Id string) (*Wave, error) {
	wr := new(waveResponse)
	endpoint := fmt.Sprintf("info/waves/byId/%s", Id)
	err := l.GetRequest(endpoint, nil, wr)
	if err != nil {
		return nil, err
	}
	wave, err := newWave(wr)
	if err != nil {
		return nil, err
	}
	return wave, nil
}
