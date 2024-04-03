package handlers

import (
	"github.com/gin-gonic/gin"
	"parking-app-mobile-api-gateway/internal/services"
)

type CarHandler struct {
	CarService *services.CarService
}

func NewCarHandler() *CarHandler {
	return &CarHandler{
		CarService: services.NewCarService(),
	}
}

// GetCars godoc
// @Tags cars
// @Summary Get all cars
// @Description Get all cars
// @Produce json
// @Success 200 {array} services.Car
// @Router /cars [get]
func (h *CarHandler) GetCars(c *gin.Context) {
	cars, err := h.CarService.GetCars()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cars)
}

// CreateCar godoc
// @Tags cars
// @Summary Create a car
// @Description Create a car
// @Accept json
// @Produce json
// @Param car body services.CreateCarRequest true "Car"
// @Success 200 {string} string
// @Router /cars [post]
func (h *CarHandler) CreateCar(c *gin.Context) {

	var createCarRequest services.CreateCarRequest
	err := c.ShouldBindJSON(&createCarRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.CarService.CreateCar(createCarRequest)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Car created",
	})
}

// GetCar godoc
// @Tags cars
// @Summary Get a car
// @Description Get a car
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} services.Car
// @Router /cars/{id} [get]
func (h *CarHandler) GetCar(c *gin.Context) {
	id := c.Param("id")

	car, err := h.CarService.GetCarByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(200, car)
}

// DeleteCar godoc
// @Tags cars
// @Summary Delete a car
// @Description Delete a car
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {string} string
// @Router /cars/{id} [delete]
func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	err := h.CarService.DeleteCarByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Car deleted",
	})
}
