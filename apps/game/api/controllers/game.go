package controllers

import (
	"blackgo/engine/exceptions"
	"blackgo/game/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New Game godoc
// @Summary  Create new game
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  201
// @Router   /game [post]
func NewGame(c *gin.Context) {
	game := repository.CreateGame()
	if game == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "error_creating_game",
		})
		return
	}
	c.JSON(http.StatusCreated, game)
}

// Open Game godoc
// @Description  Get game by ID
// @Param        id   path      string  true  "Game ID"
// @Summary  See game details
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /game/{id} [get]
func GameDetail(c *gin.Context) {
	id := c.Param("id")
	game := repository.GetGameById(id)
	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}

	c.JSON(http.StatusOK, game)
}

// Ask for another card godoc
// @Description  Ask for another card
// @Param        id   path      string  true  "Game ID"
// @Summary  Ask for another card
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Failure      400  {object}  exceptions.HTTPError
// @Router   /game/{id}/hit [post]
func Hit(c *gin.Context) {
	id := c.Param("id")
	game := repository.GetGameById(id)
	if game == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}

	err := game.Hit()
	if err != nil {
		if err == exceptions.ErrGameIsOver {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	repository.SaveGame(game)
	c.JSON(http.StatusOK, game)
}

// Stand for another card godoc
// @Description  Stand
// @Param        id   path      string  true  "Game ID"
// @Summary  Stand
// @Tags     blackgo
// @Accept   json
// @Produce  json
// @Success  200
// @Router   /game/{id}/stand [post]
func Stand(c *gin.Context) {
	id := c.Param("id")
	game := repository.GetGameById(id)
	if game == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "page_not_found",
			"message": "Page not found",
		})
		return
	}
	game.Stand()
	repository.SaveGame(game)
	c.JSON(http.StatusOK, game)
}
