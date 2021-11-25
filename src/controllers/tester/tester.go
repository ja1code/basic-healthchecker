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

func TestHost(url string) (bool, *customError.HttpError) {
	fmt.Println("[INFO] Testing", url)

	response, err := http.Get(url)

	body, _ := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("[ERROR] Unable to request ", url)
	}

	if !ACCEPTABLE_HTTP_CODES[response.StatusCode] {
		errorStruct := &customError.HttpError{
			Url:        url,
			StatusCode: response.StatusCode,
			Body:       body,
		}

		return false, errorStruct
	}

	return true, nil
}
