package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfiguration loads repository configurations from a json file into a configuration model
func LoadConfiguration(configPath string) Configuration {
	// Open config file
	file, err := os.Open(configPath)
	defer file.Close()
	Check(err)

	// Read opened file as byte array
	byteValue, _ := ioutil.ReadAll(file)

	var configuration Configuration
	json.Unmarshal(byteValue, &configuration)

	return configuration
}
