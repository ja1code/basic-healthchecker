package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestData struct {
	StatusCode int
	Body       string
	Headers    string
}

func GenericLog(url string, successInd bool, statusCode int) bool {
	file, fileOpenErr := os.OpenFile("output/general.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0655)

	if fileOpenErr != nil {
		fmt.Println(fileOpenErr)
		fmt.Println("[ERROR] Unable to open log file! Make sure the folder is r/w")
		return false
	}

	file.WriteString("[LOG] " + url + " | " + successIndStr(successInd) + " | STATUSCDOE: " + strconv.Itoa(statusCode) + " | @ " + time.Now().Format(time.RFC3339) + "\n")

	file.Close()

	return true
}

func CsvLog(url string, successInd bool, statusCode int) bool {

	urlSplit := strings.Split(url, "://")
	urlSplit = strings.Split(urlSplit[1], "/")

	file, fileOpenErr := os.OpenFile("output/general.csv", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	if fileOpenErr != nil {
		fmt.Println("[ERROR] Unable to open log file! Make sure the folder is r/w")
		return false
	}

	contentByte, fileReadErr := ioutil.ReadAll(file)

	if fileReadErr != nil {
		fmt.Println("[ERROR] Unable to read log file!")
		return false
	}

	content := string(contentByte)

	if content == "" {
		content += "url,success_ind,status_code,time\n"
	}

	content += url + "," + successIndStr(successInd) + "," + strconv.Itoa(statusCode) + "," + time.Now().Format(time.RFC3339) + "\n"

	ioutil.WriteFile("output/general.csv", []byte(content), 0666)

	return true
}

func SpecificLog(url string, data RequestData) bool {

	urlSplit := strings.Split(url, "://")
	urlSplit = strings.Split(urlSplit[1], "/")

	file, fileOpenErr := os.OpenFile("output/"+urlSplit[0]+".log.json", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	if fileOpenErr != nil {
		fmt.Println("[ERROR] Unable to open log file! Make sure the folder is r/w")
		return false
	}

	contentByte, fileReadErr := ioutil.ReadAll(file)

	if fileReadErr != nil {
		fmt.Println("[ERROR] Unable to read log file!")
		return false
	}

	var liveContent []RequestData

	if string(contentByte) != "" {
		json.Unmarshal(contentByte, &liveContent)
	}

	liveContent = append(liveContent, data)

	liveContentJson, parseError := json.Marshal(liveContent)

	if parseError != nil {
		return false
	}

	ioutil.WriteFile("output/"+urlSplit[0]+".log.json", liveContentJson, 0666)

	return true
}

func successIndStr(success_ind bool) string {
	if success_ind {
		return "SUCCESSFUL"
	} else {
		return "FAILED"
	}
}
