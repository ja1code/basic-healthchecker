package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var ALLOWED_TOKEN_POSITIONS map[string]bool = map[string]bool{ // Used as a ENUM to validate either the position specified can or cannot be used w/ the app
	"HEADER": true,
	"BODY":   true, // Assumes that the expected body is in JSON format | TODO: Add other body types support
}

var ALERT_DESTINATIONS map[string]bool = map[string]bool{ // Used as a ENUM to validade either the Alert destination is or isn't currently supported | Could be done w/ a file read or httpRequest (dynamic available destination fetch)
	"DISCORD":  true,
	"SLACK":    true,
	"TELEGRAM": true,
}

type AuthDesc struct { // Describes basic Auth
	Token    string `mapstructure:"token"`
	Position string `mapstructure:"position"`
	Key      string `mapstructure:"key"`
	Prefix   string `mapstructure:"prefix"`
}

type GenericAlertDestination struct {
	Token   string `mapstructure:"token"`
	Channel string `mapstructure:"channel"` // Discord channel Id, Slack Channel Id, Telegram Group/Channel Id
}

type ObservedHost struct {
	Url               string                    `mapstructure:"url"`
	Auth              AuthDesc                  `mapstructure:"auth"`
	Priority          int                       `mapstructure="priority"`
	AlertDestinations []GenericAlertDestination `mapstructure="destinations"`
}

var config []ObservedHost

func loadConfig() []ObservedHost {
	var readConfigs []ObservedHost

	rawFile, fileOpenErr := os.Open("config.json")

	if fileOpenErr != nil {
		fmt.Printf("[ERROR] Unable to find/open config file!\nMake sure it exists and is accessible.")
		os.Exit(0)
	}

	byteValue, fileReadErr := ioutil.ReadAll(rawFile)

	if fileReadErr != nil {
		fmt.Printf("[ERROR] Unable to read config file!\nMake sure it exists and isn't corrupted.")
		os.Exit(0)
	}

	json.Unmarshal(byteValue, &readConfigs)

	rawFile.Close()

	return readConfigs
}
