package handlers

import (
	"github.com/gin-gonic/gin"
	payments "parking-app-mobile-api-gateway/internal/payments/proto"
	"parking-app-mobile-api-gateway/internal/services"
)

type PaymentHandler struct {
	PaymentService *services.PaymentService
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		PaymentService: services.NewPaymentService(),
	}
}

// GetPayments godoc
// @Tags payments
// @Summary Get all payments
// @Description Get all payments
// @Produce json
// @Success 200 {array} payments.GetAllPaymentsResponse
// @Router /payments [get]
func (p *PaymentHandler) GetPayments(c *gin.Context) {
	payments := p.PaymentService.GetPayments()
	c.JSON(200, payments)
}

// CreatePayment godoc
// @Tags payments
// @Summary Create a payment
// @Description Create a payment
// @Accept json
// @Produce json
// @Param payment body payments.CreatePaymentRequest true "Payment"
// @Success 200 {string} string
// @Router /payments [post]
func (p *PaymentHandler) CreatePayment(c *gin.Context) {

	var createPaymentRequest payments.CreatePaymentRequest
	err := c.ShouldBindJSON(&createPaymentRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	payment := p.PaymentService.CreatePayment(&createPaymentRequest)
	c.JSON(200, payment)
}

// GetPayment godoc
// @Tags payments
// @Summary Get a payment
// @Description Get a payment
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} payments.GetPaymentResponse
// @Router /payments/{id} [get]
func (p *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment := p.PaymentService.GetPayment(id)
	c.JSON(200, payment)
}
