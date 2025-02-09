package controllers

import (
	"cars/app/config"
	"cars/app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func DeleteCarVariant(context *gin.Context) {
	carId := cast.ToUint(context.Param("car_id"))
	carVariantId := cast.ToUint(context.Param("car_variant_id"))

	carVariant := models.CarVariant{}
	res := config.DB.Where("id = ?", carVariantId).First(&carVariant)
	if res.Error != nil || carVariant.CarID != carId {
		context.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}

	// Hard delete with Unscoped
	delete := config.DB.Where("id = ?", carVariantId).Unscoped().Delete(&carVariant)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    carVariantId,
	})
}
