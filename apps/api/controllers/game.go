package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// New Game godoc
// @Summary  Create new game
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  308
// @Router   /api/v1/game/new [get]
func NewGame(c *gin.Context) {
	url := "/api/v1/game/" + uuid.NewString()
	fmt.Println(url)
	c.Redirect(http.StatusPermanentRedirect, url)
}

// Open Game godoc
// @Summary  See game details
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /api/v1/game/:id [get]
func GameDetail(c *gin.Context) {
	id := c.Param("id")
	game := games[id]
	if game == nil {
		fmt.Println("New game created at", id)
		newGame := CreateGame()
		games[id] = &newGame
		game = &newGame

		game.Start()
	}

	c.JSON(http.StatusOK, gin.H(game.JSON()))
}
