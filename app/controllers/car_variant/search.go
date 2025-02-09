package controllers

import (
	"net/http"

	"cars/app/config"
	"cars/app/models"

	"github.com/gin-gonic/gin"
)

func SearchCarVariants(context *gin.Context) {
	var car_variants []models.CarVariant

	err := config.DB.Find(&car_variants)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    car_variants,
	})
}
