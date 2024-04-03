package routes

import (
	"github.com/gin-gonic/gin"
	"parking-app-mobile-api-gateway/internal/handlers"
)

func setupPaymentRoutes(app *gin.Engine) {

	paymentHandler := handlers.NewPaymentHandler()

	app.GET("/payments", paymentHandler.GetPayments)
	app.GET("/payments/:id", paymentHandler.GetPayment)
	app.POST("/payments", paymentHandler.CreatePayment)
}
