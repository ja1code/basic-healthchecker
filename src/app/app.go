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
		_, error := tester.TestHost(config.Url)

		println(&error) // fkn address
		if error != nil {
			println("[ERROR] Error while testing", config.Url)
			println(error) // fkn address
		}
	}

}
