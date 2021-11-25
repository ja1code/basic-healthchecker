package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	Type    string `mapstructuere:"type"`
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

func LoadConfig() []ObservedHost {
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

	for _, config := range readConfigs {
		validateAuthFields(config)

		validateAlertDestinations(config.AlertDestinations)
	}

	fmt.Println("[INFO] Config file loaded and validated successfully!")
	return readConfigs
}

func validateAuthFields(config ObservedHost) {
	if config.Auth.Token == "" {
		return
	}

	config.Auth.Position = strings.ToUpper(config.Auth.Position)

	if config.Auth.Position == "" {
		config.Auth.Position = "HEADER"
	}

	if config.Auth.Prefix == "" {
		config.Auth.Prefix = "Basic"
	}

	if config.Auth.Key == "" {
		config.Auth.Key = "Authorization"
	}

	if !ALLOWED_TOKEN_POSITIONS[config.Auth.Position] {
		fmt.Println("[ERROR] Auth token position", config.Auth.Position, "isn't recognized.")
		os.Exit(0)
	}
}

func validateAlertDestinations(destinations []GenericAlertDestination) {
	if len(destinations) == 0 {
		fmt.Println("[WARNING] No alert destinations found, only logs will be created.")
	}

	for _, destination := range destinations {
		destination.Type = strings.ToUpper(destination.Type)
		if !ALERT_DESTINATIONS[destination.Type] {
			fmt.Println("[ERROR] Invalid/Inactive destination", destination, ", check the README.md file for all available destinations.")
			os.Exit(0)
		}
	}
}
