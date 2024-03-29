package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*****************
	  Setup
 *****************/

func init() {
	gin.SetMode(gin.TestMode)
}

/*****************
	  Tests
 *****************/

func TestServer_HelloWorld(T *testing.T) {
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
		NewServer().ServeHTTP(rec, req)

		if t.expectedStatusCode != rec.Code {
			T.Errorf("expected a status code of %d, but got %d", t.expectedStatusCode, rec.Code)
			T.Errorf("msg: %s", rec.Body.String())
		}
	}
}
