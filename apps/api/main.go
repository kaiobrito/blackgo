package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


// @title           Swagger Example API

func main() {

	r := setupRouter()
	r.SetTrustedProxies([]string{"localhost"})
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/game/new", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/game/"+uuid.New().String())
	})

	r.GET("/game/:id", func(c *gin.Context) {
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
	})

	return r
}
