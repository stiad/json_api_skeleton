package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stiad/json_api_skeleton/src/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ApiKeyAuth(T *testing.T) {
	server := setup()

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

func setup() *app.Server {
	gin.SetMode(gin.TestMode)
	viper.Set("API_KEY", "test-key")
	server := app.NewServer()
	return server
}
