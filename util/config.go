package util

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStructure struct {
	ServerPort  int
	ApiEndpoint string
    DatabaseAddress string
}

var Config ConfigStructure

func ReadConfig(filePath string) {
    Logger.Printf("Reading config from %s", filePath)
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		Logger.Fatalf("Error while reading config file: %s", err)
	}
	err = json.Unmarshal(rawData, &Config)
	if err != nil {
		Logger.Fatalf("Error while decoding json input in config file: %s", err)
	}
}
