package controllers

import (
	"cars/app/config"
	"cars/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func ShowCar(context *gin.Context) {
	carId := cast.ToUint(context.Param("car_id"))
	car := models.Car{}

	res := config.DB.Where("id = ?", carId).First(&car)
	if res.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failure", "error": "car not found"})
		return
	}

	var car_variants []models.CarVariant

	err := config.DB.Find(&car_variants)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}
	car.CarVariants = car_variants

	context.JSON(http.StatusOK, car)
}
