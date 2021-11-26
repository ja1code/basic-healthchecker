package tester

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ja1code/basic-healthchecker/src/entities/customError"
)

var ACCEPTABLE_HTTP_CODES map[int]bool = map[int]bool{
	200: true,
	201: true,
	202: true,
	203: true,
	204: true,
}

type SuccessfulResponseData struct {
	StatusCode int
	Body       string
	Header     map[string][]string // Header type copied from http src
}

func TestHost(url string) (*SuccessfulResponseData, *customError.HttpError) {
	fmt.Println("[INFO] Testing", url)

	response, err := http.Get(url)

	body, _ := io.ReadAll(response.Body)
	header := response.Header

	if err != nil {
		fmt.Println("[ERROR] Unable to request ", url)
	}

	if !ACCEPTABLE_HTTP_CODES[response.StatusCode] {
		errorStruct := &customError.HttpError{
			Url:        url,
			StatusCode: response.StatusCode,
			Body:       body,
			Header:     header,
		}

		return nil, errorStruct
	}

	successData := &SuccessfulResponseData{
		StatusCode: response.StatusCode,
		Body:       string(body),
		Header:     header,
	}

	return successData, nil
}
