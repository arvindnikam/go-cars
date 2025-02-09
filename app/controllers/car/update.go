package controllers

import (
	"cars/app/config"
	"cars/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oleiade/reflections"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

func UpdateCar(context *gin.Context) {
	carId := cast.ToUint(context.Param("car_id"))

	// logger with defualt fields to logger
	logger := logrus.WithFields(logrus.Fields{
		"car_id":       carId,
		"path":         context.Request.URL.Path,
		"Content-Type": context.Request.Header["Content-Type"],
	})

	car := models.Car{}
	res := config.DB.Where("id = ?", carId).First(&car)
	if res.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failure", "error": "car not found"})
		return
	}

	var data carRequest
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.WithField("data", data).Printf("Car update request received: %#v", data)

	dataKeys := []string{"Make", "CarModel", "Year", "BodyType"}
	for _, key := range dataKeys {
		if value, err := reflections.GetField(data, key); err == nil && value != "" && value != 0 {
			reflections.SetField(&car, key, value)
		}
	}

	result := config.DB.Save(&car)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// logger.WithField("updated_car", car).Info("Car updated successfully")

	context.JSON(http.StatusOK, car)
}
