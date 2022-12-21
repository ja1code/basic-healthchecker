package app

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ja1code/basic-healthchecker/src/entities/config"
	"github.com/ja1code/basic-healthchecker/src/services/log"
	"github.com/ja1code/basic-healthchecker/src/services/tester"
)

func StartApp() {
	var configs []config.ObservedHost = config.LoadConfig()
	os.Mkdir("output", 0777) // Placed here for the moment, when proper first time setup exists it will be placed there

	if len(configs) == 0 {
		println("[ERROR] No config found, exiting process")
		os.Exit(0)
	}

	for {
		MainRoutine(configs)
		time.Sleep(20 * time.Second)
	}

}

func MainRoutine(configs []config.ObservedHost) {
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
			log.CsvLog(config.Url, false, err.StatusCode)

			println("[ERROR] Error while testing", config.Url)
			continue
		}

		log.CsvLog(config.Url, true, successData.StatusCode)
		log.GenericLog(config.Url, true, successData.StatusCode)
	}
}
