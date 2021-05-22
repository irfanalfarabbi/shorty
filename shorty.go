package main

import (
  "fmt"
  "log"
  "net/http"
  "time"

  "irfanalfarabbi/shorty/handler"
  "irfanalfarabbi/shorty/logger"

  "github.com/julienschmidt/httprouter"
  "github.com/rs/cors"
)

func main() {
  router := httprouter.New()
  router.GET("/healthz", healthz)
  router.POST("/api/shorten", handler.Shorten)
  router.GET("/api/:shortcode", handler.Shortcode)
  router.GET("/api/:shortcode/stats", handler.ShortcodeStats)

  co := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST"},
    AllowedHeaders: []string{"*"},
    MaxAge:         86400,
  })

  server := &http.Server{
    Addr:         ":8080",
    Handler:      co.Handler(router),
    ReadTimeout:  300 * time.Second,
    WriteTimeout: 300 * time.Second,
  }

  fmt.Println("Server is running at :8080")
  log.Fatal(server.ListenAndServe())
}

func healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var method = http.MethodGet
  var statusCode = http.StatusOK
  var response = "OK"

  w.WriteHeader(statusCode)
  w.Write([]byte(response))
  logger.Request(method, "/healthz", statusCode, response)
}
