package handler

import (
  "encoding/json"
  "net/http"
  "time"

  "irfanalfarabbi/shorty/logger"
  "irfanalfarabbi/shorty/service"

  "github.com/julienschmidt/httprouter"
)

const (
  SHORTCODE_API_URL      = "/:shortcode"
  SHORTCODESTATS_API_URL = "/:shortcode/stats"
  TIME_FORMAT            = time.RFC3339
)

type shortcodeStatsResponse struct {
  StartDate     string `json:"startDate"`
  LastSeenDate  string `json:"lastSeenDate,omitempty"`
  RedirectCount int    `json:"redirectCount"`
}

func Shortcode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  var statusCode = http.StatusFound
  var shortenURL *service.ShortenURL

  shortcode := params.ByName("shortcode")

  if shortcode != "" {
    shortenURL = service.GetShortenURL(shortcode, true)
  }

  if shortenURL == nil {
    statusCode = http.StatusNotFound
    logger.Request(http.MethodGet, SHORTCODE_API_URL, statusCode, nil)
    w.WriteHeader(statusCode)
    return
  }

  shortenURL.RedirectCount++

  logger.Request(http.MethodGet, SHORTCODE_API_URL, statusCode, shortenURL.URL)
  w.Header().Set("Location", shortenURL.URL)
  w.WriteHeader(statusCode)
  json.NewEncoder(w).Encode("")
}

func ShortcodeStats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  var statusCode = http.StatusOK
  var shortenURL *service.ShortenURL
  var response shortcodeStatsResponse

  shortcode := params.ByName("shortcode")

  if shortcode != "" {
    shortenURL = service.GetShortenURL(shortcode, false)
  }

  if shortenURL == nil {
    statusCode = http.StatusNotFound
    logger.Request(http.MethodGet, SHORTCODESTATS_API_URL, statusCode, nil)
    w.WriteHeader(statusCode)
    return
  }

  response.StartDate = shortenURL.StartDate.UTC().Format(TIME_FORMAT)
  response.RedirectCount = shortenURL.RedirectCount

  if !shortenURL.LastSeenDate.IsZero() {
    response.LastSeenDate = shortenURL.LastSeenDate.UTC().Format(TIME_FORMAT)
  }

  logger.Request(http.MethodGet, SHORTCODESTATS_API_URL, statusCode, shortenURL.URL)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
  json.NewEncoder(w).Encode(response)
}
