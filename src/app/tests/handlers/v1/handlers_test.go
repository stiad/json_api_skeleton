package v1

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
		expectedStatusCode int
	}{
		{http.StatusOK},
	}

	for _, t := range tt {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/hello", nil)
		server.ServeHTTP(rec, req)

		if t.expectedStatusCode != rec.Code {
			T.Errorf("expected a status code of %d, but got %d", t.expectedStatusCode, rec.Code)
		}
	}
}

func setup() *app.Server {
	gin.SetMode(gin.TestMode)
	viper.Set("API_KEY", "test-key")
	server := app.NewServer()
	return server
}
