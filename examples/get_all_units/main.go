package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("{secretKey}", "https://apiv2.legiontd2.com/")
	outFolder := "./units/"
	constantsFile := "units.txt"
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
		go func(unit string, api *ltdsdk.LegionTDSdk) {
			defer wg.Done()
			u, err := api.GetUnit(unit)
			if err != nil {
				panic(err)
			}
			u.ExportToJson("./units/")
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
