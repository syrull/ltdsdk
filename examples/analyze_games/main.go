package main

import (
	"fmt"
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
	api := ltdsdk.NewLTDSDK("bwTJFWPtIE7pa1u9KQ6Sz75gwuWPqV1o4snOzafJ", "https://apiv2.legiontd2.com/")
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
			fmt.Printf("Fetching offset: %v\n", offset)
			defer wg.Done()
		}(api, i)
	}
	wg.Wait()
}
