package handler

import (
	"net/http"

	"irfanalfarabbi/shorty/logger"
)

func returnNotFound(w http.ResponseWriter, method string, url string) {
	statusCode := http.StatusNotFound
	logger.Request(method, url, statusCode, nil)
	w.WriteHeader(statusCode)
	w.Write([]byte("404 page not found\n"))
}
