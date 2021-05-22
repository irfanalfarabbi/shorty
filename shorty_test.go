package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/julienschmidt/httprouter"
  "github.com/stretchr/testify/assert"
)

func Test_healthz(t *testing.T) {
  tests := []struct {
    name         string
    wantCode     int
    wantResponse string
  }{
    {"success", http.StatusOK, "OK"},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      req, err := http.NewRequest("GET", "/healthz", nil)
      if err != nil {
        t.Fatal(err)
      }

      rr := httptest.NewRecorder()
      router := httprouter.New()
      router.GET("/healthz", healthz)
      router.ServeHTTP(rr, req)

      assert.Equal(t, tt.wantCode, rr.Code)
      assert.Equal(t, tt.wantResponse, rr.Body.String())
    })
  }
}
