package util

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStructure struct {
	ServerHostname  string
	ServerPort      int
	ApiEndpoint     string
	DatabaseAddress string
	DatabaseName    string
}

var Config ConfigStructure

// Gets config values from the config file
func ReadConfig(filePath string) {
	Logger.Printf("Reading config from %s", filePath)

	// Read whole file in the string
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		Logger.Fatalf("Error while reading config file: %s", err)
	}
	// Get configuration values
	err = json.Unmarshal(rawData, &Config)
	if err != nil {
		Logger.Fatalf("Error while decoding json input in config file: %s", err)
	}
}
