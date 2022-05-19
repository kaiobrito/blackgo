package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func perfomRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	router.ServeHTTP(w, req)
	gin.DisableConsoleColor()
	gin.SetMode(gin.TestMode)
	return w
}

func TestPingEndpoint(t *testing.T) {
	w := perfomRequest("GET", "/ping", nil)

	assert.Equal(t, w.Code, http.StatusOK)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	fmt.Println(response)
	assert.Nil(t, err)

	if !reflect.DeepEqual(response, map[string]string{
		"message": "pong",
	}) {
		t.Errorf("Not pong")
	}

}
