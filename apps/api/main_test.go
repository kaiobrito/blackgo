package main

import (
	"blackgo/api/controllers"
	"blackgo/engine"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func perfomRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	router.ServeHTTP(w, req)
	gin.DisableConsoleColor()
	return w
}

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
		for k := range controllers.Games {
			delete(controllers.Games, k)
		}
	}
}

func TestPingEndpoint(t *testing.T) {
	w := perfomRequest("GET", "/ping", nil)

	assert.Equal(t, w.Code, http.StatusOK)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, response, map[string]string{
		"message": "pong",
	})

}

func TestNewGameEndpoint(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	w := perfomRequest("GET", "/api/v1/game", nil)
	assert.Equal(t, w.Code, http.StatusPermanentRedirect)
	assert.Regexp(t, "/api/v1/game/*", w.Header().Get("Location"))
}

func TestCreateNewGame(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	assert.Equal(t, len(controllers.Games), 0)
	w := perfomRequest("GET", "/api/v1/game/1231", nil)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(controllers.Games), 1)

	game := controllers.Games["1231"]
	assert.NotNil(t, game)
	assert.NotNil(t, game.UserDeck)
}

func TestHitUnknownGame(t *testing.T) {
	w := perfomRequest("GET", "/api/v1/game/1231/hit", nil)
	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestHit(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	// Create the game
	game := engine.NewBlackgoGame()
	game.Start()
	controllers.Games["1231"] = &game
	assert.Equal(t, len(game.UserDeck), 2)

	w := perfomRequest("GET", "/api/v1/game/1231/hit", nil)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(game.UserDeck), 3)
}
