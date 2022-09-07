package main

import (
	"sync"

	"github.com/syrull/ltdsdk"
)

const (
	DbName    = "games.sqlite"
	Version   = "v9.07.2"
	QueueType = "Normal"
	Pages     = 60
)

func main() {
	var wg sync.WaitGroup
	api := ltdsdk.NewLTDSDK("{secret_key}", "https://apiv2.legiontd2.com/")
	for i := 0; i <= Pages*50; i += 50 {
		wg.Add(1)
		go func(api *ltdsdk.LegionTDSdk, offset int) {
			gameCollection, _ := api.GetGames(&ltdsdk.GameOptions{
				Version:       Version,
				Limit:         50,
				Offset:        offset,
				SortDirection: -1,
				QueueType:     QueueType,
			})
			gameCollection.ExportToSql(DbName)
			defer wg.Done()
		}(api, i)
	}
	wg.Wait()
}
