package handler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"irfanalfarabbi/shorty/service"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestShortcode(t *testing.T) {
	service.CreateShortenURL("100001", "http://test1")
	service.CreateShortenURL("100002", "http://test2")
	service.CreateShortenURL("100003", "http://test3")

	tests := []struct {
		name         string
		shortcode    string
		wantCode     int
		wantResponse string
	}{
		{"Success", "100001", http.StatusFound, "\"\"\n"},
		{"Success", "100002", http.StatusFound, "\"\"\n"},
		{"Success", "100003", http.StatusFound, "\"\"\n"},
		{"Faield", "1", http.StatusNotFound, ""},
		{"Faield", "2", http.StatusNotFound, ""},
		{"Faield", "3", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortcodeUrl := "/api/" + tt.shortcode
			req, err := http.NewRequest(http.MethodGet, shortcodeUrl, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := httprouter.New()
			router.GET("/api/:shortcode", Shortcode)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantResponse, rr.Body.String())
		})
	}
}

func TestShortcodeStats(t *testing.T) {
	service.CreateShortenURL("100001", "http://test1")
	service.CreateShortenURL("100002", "http://test2")
	service.CreateShortenURL("100003", "http://test3")

	tests := []struct {
		name         string
		shortcode    string
		wantCode     int
		wantResponse string
	}{
		{"Success", "100001", http.StatusOK, `{"startDate":"(.+)T(.+)Z","lastSeenDate":"(.+)T(.+)Z","redirectCount":1}\n`},
		{"Success", "100002", http.StatusOK, `{"startDate":"(.+)T(.+)Z","lastSeenDate":"(.+)T(.+)Z","redirectCount":1}\n`},
		{"Success", "100003", http.StatusOK, `{"startDate":"(.+)T(.+)Z","lastSeenDate":"(.+)T(.+)Z","redirectCount":1}\n`},
		{"Faield", "1", http.StatusNotFound, ""},
		{"Faield", "2", http.StatusNotFound, ""},
		{"Faield", "3", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestURL := "/api/" + tt.shortcode + "/stats"
			req, err := http.NewRequest(http.MethodGet, requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := httprouter.New()
			router.GET("/api/:shortcode/stats", ShortcodeStats)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Regexp(t, regexp.MustCompile(tt.wantResponse), rr.Body.String())
		})
	}
}
