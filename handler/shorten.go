package handler

import (
  "encoding/json"
  "net/http"

  "irfanalfarabbi/shorty/logger"
  "irfanalfarabbi/shorty/service"

  "github.com/julienschmidt/httprouter"
)

const (
  SHORTEN_API_URL         = "/shorten"
  SHORTEN_URL_MEMOIZATION = false
)

type shortenRequest struct {
  URL       string `json:"url"`
  Shortcode string `json:"shortcode"`
}

type shortenResponse struct {
  Shortcode string `json:"shortcode"`
}

func Shorten(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var statusCode = http.StatusCreated
  var request shortenRequest
  var response shortenResponse

  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    statusCode = http.StatusInternalServerError
    logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, err)
    w.WriteHeader(statusCode)
    return
  }

  if !service.IsValidURL(request.URL) {
    statusCode = http.StatusBadRequest
    logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, nil)
    w.WriteHeader(statusCode)
    return
  }

  if SHORTEN_URL_MEMOIZATION {
    if response.Shortcode = service.GetRegisteredURL(request.URL); response.Shortcode != "" {
      logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, response)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(statusCode)
      json.NewEncoder(w).Encode(response)
      return
    }
  }

  if len(request.Shortcode) > 0 {
    if !service.IsValidShortcode(request.Shortcode) {
      statusCode = http.StatusUnprocessableEntity
      logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, nil)
      w.WriteHeader(statusCode)
      return
    }

    if service.IsShortenURLExists(request.Shortcode) {
      statusCode = http.StatusConflict
      logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, nil)
      w.WriteHeader(statusCode)
      return
    }
  } else {
    request.Shortcode = service.GetNextShortcode()
  }

  if err := service.CreateShortenURL(request.Shortcode, request.URL); err != nil {
    statusCode = http.StatusInternalServerError
    logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, err)
    w.WriteHeader(statusCode)
    return
  }

  if SHORTEN_URL_MEMOIZATION {
    service.RegisterURL(request.Shortcode, request.URL)
  }

  response.Shortcode = request.Shortcode

  logger.Request(http.MethodPost, SHORTEN_API_URL, statusCode, response)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
  json.NewEncoder(w).Encode(response)
}
