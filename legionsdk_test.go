package ltdsdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func LoadFixture(fileName string) any {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var fixture any
	json.Unmarshal([]byte(byteValue), &fixture)
	return fixture
}
