package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	controllers "blackgo/api/controllers"

	docs "blackgo/api/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	r := setupRouter()
	r.SetTrustedProxies([]string{"localhost"})
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/api/v1")
	{
		game := v1.Group("/game")
		{

			game.POST("", controllers.NewGame)
			game.GET(":id", controllers.GameDetail)
			game.POST(":id/hit", controllers.Hit)
			game.POST(":id/stand", controllers.Stand)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
