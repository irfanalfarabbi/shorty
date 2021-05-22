package logger

import (
	"encoding/json"
	"fmt"
)

func Request(method string, url string, statusCode int, response interface{}) {
	strResponse, _ := json.Marshal(response)
	fmt.Printf("Request URL: %s %s - StatusCode: %v, Response: %v\n", method, url, statusCode, string(strResponse))
}
