package main

import (
  "fmt"
  "net/http"

  "irfanalfarabbi/shorty/api"
)

func main() {
  // http.HandleFunc("/shorten", api.Shorten)
  // http.HandleFunc("/:shortcode/stats", api.ShortcodeStats)
  // http.HandleFunc("/:shortcode", api.Shortcode)

  http.HandleFunc("/healthz", healthz)

  fmt.Println("Server is running at :8080")
  http.ListenAndServe(":8080", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
  var method = http.MethodGet
  var statusCode = http.StatusOK
  var response = "OK"

  w.WriteHeader(statusCode)
  w.Write([]byte(response))
  api.LogRequest(method, "/shorten", statusCode, response)
}
