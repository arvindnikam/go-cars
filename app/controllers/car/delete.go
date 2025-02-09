package controllers

import (
	"cars/app/config"
	"cars/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"
)

func DeleteCar(context *gin.Context) {
	carId := cast.ToUint(context.Param("car_id"))

	// Initiate Car models
	car := models.Car{}
	res := config.DB.Where("id = ?", carId).First(&car)
	if res.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failure", "error": "car not found"})
		return
	}

	// Soft delete as car has DeletedAt gorm.DeletedAt
	config.DB.Where("id = ?", carId).Delete(&car)

	// Delete a car's CarVariants and CarParts associations when deleting the user, CarParts is just an example to note multiple associations
	// db.Select("CarVariants", "CarParts").Delete(&car)

	// Delete all of a car's has one, has many, and many2many associations
	config.DB.Select(clause.Associations).Delete(&car)

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    carId,
	})
}
