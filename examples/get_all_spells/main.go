package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("{secretKey}", "https://apiv2.legiontd2.com/")
	outFolder := "./spells/"
	constantsFile := "spells.txt"
	if _, err := os.Stat(outFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(outFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	var wg sync.WaitGroup
	file, err := os.Open(constantsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		go func(spellId string, api *ltdsdk.LegionTDSdk) {
			defer wg.Done()
			s, err := api.GetSpell(spellId)
			if err != nil {
				panic(err)
			}
			f, err := os.Create(outFolder + "/" + spellId + ".json")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			enc := json.NewEncoder(f)
			err = enc.Encode(s)
			if err != nil {
				panic(err)
			}
		}(scanner.Text(), api)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
