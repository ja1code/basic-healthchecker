package app

import (
	"fmt"
	"os"
	"io/ioutil"
)

func StartApp() {

}

func ReadConfig() {
	var configs

	file, openFileErr := os.Open("config.json")

	if openFileErr != nil {
		fmt.Printf("[ERROR] Unable to find/open config file!\nMake sure it exists and follows the 'config.template.json' structure.")
		os.Exit(0)
	}

	byteValue, readFileErr := ioutil.ReadAll(file)

	if readFileErr != nil {
		fmt.Printf("[ERROR] Unable to read config file!\nMake sure it exists and isn't corrupted.")
		os.Exit(0)
	}


}