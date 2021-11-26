package log

import (
	"fmt"
	"os"
	"time"
)

func GenericLog(url string, success_ind bool, statusCode int) bool {
	file, fileOpenErr := os.OpenFile("output/general.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 655)

	if fileOpenErr != nil {
		fmt.Println("[ERROR] Unable to open log file! Make sure the folder is r/w")
		os.Exit(1)
	}

	file.WriteString("[LOG] " + url + " | " + successIndStr(success_ind) + " | STATUSCDOE: " + string(statusCode) + " | @ " + time.Now().Format("stdZeroDay/"))
}

func successIndStr(success_ind bool) string {
	if success_ind {
		return "SUCCESSFUL"
	} else {
		return "FAILED"
	}
}
