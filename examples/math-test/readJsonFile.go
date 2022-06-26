package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func readJsonFile(filename string) []TestItem {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var testItems TestItems

	json.Unmarshal(byteValue, &testItems)

	return testItems.TestItems
}
