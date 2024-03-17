package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stiad/json_api_skeleton/src/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*****************
	 Tests
 *****************/

func TestServer_HandlerHelloWorld(T *testing.T) {
	server := setup()

	tt := []struct {
		url                string
		expectedStatusCode int
	}{
		{"/v1/hello", http.StatusOK},
		{"/v1/fake", http.StatusNotFound},
	}

	for _, t := range tt {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", t.url, nil)
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
	return app.NewServer()
}
