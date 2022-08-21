package main

import (
	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("{secret_key}", "https://apiv2.legiontd2.com/")
	gameCollection, _ := api.GetGames(&ltdsdk.GameOptions{})
	gameCollection.ExportToJson("games.json")
}
