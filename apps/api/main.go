package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	return r
}
