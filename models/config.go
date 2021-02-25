package models

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"strings"
	"os"
)

type ServerConfig struct {
	Homeserver string
}

type ConfigCollection map[string]ServerConfig


func LoadConfig() (*ConfigCollection,error) {
	
	configFilePath := GetPathForFile("config.json")
	cfgCollection := make(ConfigCollection)

	if _,err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFile,_ := json.MarshalIndent(cfgCollection,""," ")
		_ = ioutil.WriteFile(configFilePath,configFile,0744)
	}

	jsonStream,err := ioutil.ReadFile(configFilePath)
	if (err != nil) {
		return nil,err
	}

	dec := json.NewDecoder(strings.NewReader(string(jsonStream)))

	for {

		if err := dec.Decode(&cfgCollection);err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

	}

	return &cfgCollection,nil;
}