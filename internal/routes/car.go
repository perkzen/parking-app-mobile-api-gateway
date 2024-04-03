package routes

import (
	"github.com/gin-gonic/gin"
	"parking-app-mobile-api-gateway/internal/handlers"
)

func setupCarRoutes(app *gin.Engine) {

	carHandler := handlers.NewCarHandler()

	app.GET("/cars", carHandler.GetCars)
	app.GET("/cars/:id", carHandler.GetCar)
	app.POST("/cars", carHandler.CreateCar)
	app.DELETE("/cars/:id", carHandler.DeleteCar)
}
