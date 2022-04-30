package testhelpers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func AssertResponseString(T *testing.T, router *gin.Engine, req *http.Request, expectedStatus int, expectedText string) {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(T, expectedStatus, w.Code)
	assert.Equal(T, "\""+expectedText+"\"", w.Body.String())
}

func GetResponse(T *testing.T, router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
