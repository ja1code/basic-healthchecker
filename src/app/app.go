package app

import (
	"encoding/json"
	"os"

	"github.com/ja1code/basic-healthchecker/src/controllers/log"
	"github.com/ja1code/basic-healthchecker/src/controllers/tester"
	"github.com/ja1code/basic-healthchecker/src/entities/config"
)

func StartApp() {
	var configs []config.ObservedHost = config.LoadConfig()
	os.Mkdir("output", 0777) // Placed here for the moment, when proper first time setup exists it will be placed there

	if len(configs) == 0 {
		println("[ERROR] No config found, exiting process")
		os.Exit(0)
	}

	for _, config := range configs {
		successData, err := tester.TestHost(config.Url)

		if err != nil {

			log.GenericLog(config.Url, false, err.StatusCode)

			headerJson, _ := json.Marshal(err.Header)

			requestData := log.RequestData{
				StatusCode: err.StatusCode,
				Body:       string(err.Body),
				Headers:    string(headerJson),
			}

			log.SpecificLog(config.Url, requestData)

			println("[ERROR] Error while testing", config.Url)
			continue
		}

		log.GenericLog(config.Url, true, successData.StatusCode)
	}

}
