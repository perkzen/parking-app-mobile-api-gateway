package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"parking-app-mobile-api-gateway/docs"
)

func InitRouter(app *gin.Engine) {

	docs.SwaggerInfo.Title = "Parking app Mobile API Gateway"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	app.Use(cors.New(config))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Mobile API Gateway",
		})
	})

	setupCarRoutes(app)
	setupPaymentRoutes(app)
}
