package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/arminazimi/Basic_WebServer/webportal"
)

type cofiguration struct {
	Webserver string
}

func main() {
	config := LoadConfiguration("config.json")
	webportal.Portal(config.Webserver)

}

func LoadConfiguration(file string) cofiguration {
	var config cofiguration
	configFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
