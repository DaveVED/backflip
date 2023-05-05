package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type BackflipConfiguration struct {
	TargetURL string `json:"targetUrl"`
}

type BackflipConfigurationFile struct {
	Target target `json:"target"`
}

type target struct {
	Url string `json:"url"`
}

func ParseConfigFile(filePath string) BackflipConfiguration {
	var backflipConfiguration BackflipConfiguration
	var backflipConfigurationfile BackflipConfigurationFile

	configurationFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to open configuration file %v", filePath)
		os.Exit(3)
	}

	log.Printf("Successfully found/opend %v", filePath)
	defer configurationFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(configurationFile)
	json.Unmarshal(byteValue, &backflipConfigurationfile)

	// setup configuration.
	backflipConfiguration.TargetURL = backflipConfigurationfile.Target.Url

	return backflipConfiguration
}
