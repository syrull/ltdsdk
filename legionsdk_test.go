package ltdsdk_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadFixture(fileName string) string {
	var fixture map[string]interface{}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("error when opening file: ", err)
	}
	err = json.Unmarshal(content, &fixture)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	jsonStr, err := json.Marshal(fixture)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return string(jsonStr)
}
