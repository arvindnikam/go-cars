package controllers

import (
	"net/http"

	"cars/app/config"
	"cars/app/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cast"
)

func CreateCarVariant(context *gin.Context) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	carId := cast.ToUint(context.Param("car_id"))
	car := models.Car{}
	res := config.DB.Where("id = ?", carId).First(&car)
	if res.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failure", "error": "car not found"})
		return
	}

	carVariant := models.CarVariant{}
	if err := context.ShouldBindJSON(&carVariant); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carVariant.Car = car

	result := config.DB.Create(&carVariant)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msg("")
		context.JSON(http.StatusBadRequest, gin.H{"status": "failure", "error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, carVariant)
}
