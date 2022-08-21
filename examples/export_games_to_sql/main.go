package main

import (
	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("{secretKey}", "https://apiv2.legiontd2.com/")
	gameCollection, _ := api.GetGames(&ltdsdk.GameOptions{})
	gameCollection.ExportToSql("games.sqlite")
}
