package app

import (
	"os"

	"github.com/ja1code/basic-healthchecker/src/controllers/tester"
	"github.com/ja1code/basic-healthchecker/src/entities/config"
)

func StartApp() {
	var configs []config.ObservedHost = config.LoadConfig()

	if len(configs) == 0 {
		println("[ERROR] No config found, exiting process")
		os.Exit(0)
	}

	for _, config := range configs {
		_, err := tester.TestHost(config.Url)

		if err != nil {
			println("[ERROR] Error while testing", config.Url)
			println(string(err.Body)) // fkn address
		}
	}

}
