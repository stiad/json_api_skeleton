package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*****************
	  Setup
 *****************/

func init() {
	gin.SetMode(gin.TestMode)
	viper.Set("API_KEY", "test-key")
}

/*****************
	  Tests
 *****************/

func TestServer_handleApiKeyAuth(T *testing.T) {
	server := NewServer()

	tt := []struct {
		apiKey             string
		url                string
		expectedStatusCode int
	}{
		{"test-key", "/v1/protected/hello", http.StatusOK},
		{"wrong-key", "/v1/protected/hello", http.StatusUnauthorized},
	}

	for _, t := range tt {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", t.url, nil)
		req.Header.Set("x-api-key", t.apiKey)

		server.ServeHTTP(rec, req)

		if t.expectedStatusCode != rec.Code {
			T.Errorf("expected a status code of %d, but got %d", t.expectedStatusCode, rec.Code)
			T.Errorf("msg: %s", rec.Body.String())
		}
	}
}
