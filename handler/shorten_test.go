package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name         string
		request      string
		wantCode     int
		wantResponse string
	}{
		{"Success#1", "{\"url\":\"http://test1\"}", http.StatusCreated, "{\"shortcode\":\"1_____\"}\n"},
		{"Success#2", "{\"url\":\"http://test2\"}", http.StatusCreated, "{\"shortcode\":\"2_____\"}\n"},
		{"Success#3", "{\"url\":\"http://test3\",\"shortcode\":\"123456\"}", http.StatusCreated, "{\"shortcode\":\"123456\"}\n"},
		{"Success#4", "{\"url\":\"http://test4\",\"shortcode\":\"123457\"}", http.StatusCreated, "{\"shortcode\":\"123457\"}\n"},
		{"FailedInvalidRequest", "test", http.StatusInternalServerError, ""},
		{"FailedEmptyRequest", "{}", http.StatusBadRequest, ""},
		{"FailedEmptyURL", "{\"url\":\"\"}", http.StatusBadRequest, ""},
		{"FailedInvalidURL", "{\"url\":\"test\"}", http.StatusBadRequest, ""},
		{"FailedTooShortCode#1", "{\"url\":\"http://test\",\"shortcode\":\"12345\"}", http.StatusUnprocessableEntity, ""},
		{"FailedTooLongCode#2", "{\"url\":\"http://test\",\"shortcode\":\"1234567\"}", http.StatusUnprocessableEntity, ""},
		{"FailedInvalidCode", "{\"url\":\"http://test\",\"shortcode\":\"!@#$^&\"}", http.StatusUnprocessableEntity, ""},
		{"FailedExistsCode", "{\"url\":\"http://test\",\"shortcode\":\"123456\"}", http.StatusConflict, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := bytes.NewBuffer([]byte(tt.request))
			req, err := http.NewRequest(http.MethodPost, "/api/shorten", request)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := httprouter.New()
			router.POST("/api/shorten", Shorten)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantResponse, rr.Body.String())
		})
	}
}
