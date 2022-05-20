package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// New Game godoc
// @Summary  Create new game
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  308
// @Router   /game [get]
func NewGame(c *gin.Context) {
	game := CreateGame()
	c.JSON(http.StatusCreated, gin.H(game.JSON()))
}

// Open Game godoc
// @Description  Get game by ID
// @Param        id   path      string  true  "Game ID"
// @Summary  See game details
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /game/:id [get]
func GameDetail(c *gin.Context) {
	id := c.Param("id")
	game := Games[id]
	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H(game.JSON()))
}

// Ask for another card godoc
// @Description  Ask for another card
// @Param        id   path      string  true  "Game ID"
// @Summary  Ask for another card
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /game/:id/hit [post]
func Hit(c *gin.Context) {
	id := c.Param("id")
	game := Games[id]
	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}
	game.Hit()
}

// Stand for another card godoc
// @Description  Ask for another card
// @Param        id   path      string  true  "Game ID"
// @Summary  Ask for another card
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /game/:id/stand [post]
func Stand(c *gin.Context) {
	id := c.Param("id")
	game := Games[id]
	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}
	game.Stand()
}
