package main

import (
	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("PMOR5UE0ICmBZPdwpBwi1fl22ulgmd999NVcUSN8", "https://apiv2.legiontd2.com/")
	gameCollection, _ := api.GetGames(&ltdsdk.GameOptions{})
	gameCollection.ExportToSql("games.sqlite")
}
