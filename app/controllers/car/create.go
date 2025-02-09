package controllers

import (
	"net/http"

	"cars/app/config"
	"cars/app/models"

	"github.com/gin-gonic/gin"
)

func CreateCar(context *gin.Context) {
	car := models.Car{}

	// Binding request body json to car model struct
	if err := context.ShouldBindJSON(&car); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	result := config.DB.Create(&car)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Creating http response
	context.JSON(http.StatusCreated, car)
}
