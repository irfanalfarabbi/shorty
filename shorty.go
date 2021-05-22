package main

import (
  "fmt"
  "net/http"

  "irfanalfarabbi/shorty/handler"
  "irfanalfarabbi/shorty/logger"
)

func main() {
  http.HandleFunc("/shorten", handler.Shorten)
  // http.HandleFunc("/:shortcode/stats", handler.ShortcodeStats)
  // http.HandleFunc("/:shortcode", handler.Shortcode)

  http.HandleFunc("/healthz", healthz)

  fmt.Println("Server is running at :8080")
  fmt.Println(http.ListenAndServe(":8080", nil))
}

func healthz(w http.ResponseWriter, r *http.Request) {
  var method = http.MethodGet
  var statusCode = http.StatusOK
  var response = "OK"

  w.WriteHeader(statusCode)
  w.Write([]byte(response))
  logger.Request(method, "/shorten", statusCode, response)
}
