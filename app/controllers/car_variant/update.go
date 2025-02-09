package controllers

import (
	"cars/app/config"
	"cars/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oleiade/reflections"
	"github.com/spf13/cast"
)

func UpdateCarVariant(context *gin.Context) {
	var data carVariantRequest

	carId := cast.ToUint(context.Param("car_id"))
	carVariantId := cast.ToUint(context.Param("car_variant_id"))

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carVariant := models.CarVariant{}
	res := config.DB.Where("id = ?", carVariantId).First(&carVariant)
	if res.Error != nil || carVariant.CarID != carId {
		context.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}

	dataKeys := []string{"VariantCode", "VariantName", "Transmission", "Color", "Engine"}
	for _, key := range dataKeys {
		if value, err := reflections.GetField(data, key); err == nil && value != "" && value != 0 {
			reflections.SetField(&carVariant, key, value)
		}
	}

	result := config.DB.Save(&carVariant)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, carVariant)
}
