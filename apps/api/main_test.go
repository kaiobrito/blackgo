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

func TestGetGame(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	game := controllers.CreateGame()

	w := perfomRequest("GET", "/api/v1/game/"+game.ID, nil)
	assert.Equal(t, w.Code, http.StatusOK)
	expected, _ := json.Marshal(game.JSON())
	assert.Equal(t, w.Body.String(), string(expected))
}

func TestAccessUnknownGame(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	assert.Equal(t, len(controllers.Games), 0)
	w := perfomRequest("GET", "/api/v1/game/1231", nil)
	assert.Equal(t, w.Code, http.StatusNotFound)
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

	w := perfomRequest("POST", "/api/v1/game/1231/hit", nil)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(game.UserDeck), 3)
}

func TestStand(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	// Create the game
	game := engine.NewBlackgoGame()
	game.Start()
	controllers.Games["1231"] = &game
	assert.Equal(t, game.Winner, engine.NOONE)

	w := perfomRequest("POST", "/api/v1/game/1231/stand", nil)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, game.Winner, engine.NOONE)
}
