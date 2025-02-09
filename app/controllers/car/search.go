package controllers

import (
	"fmt"
	"net/http"

	"cars/app/config"
	"cars/app/helpers"
	"cars/app/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// get all cars
func SearchCars(context *gin.Context) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var data SearchRequest
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info().Msg(fmt.Sprintf("Car update request received: %#v", data))

	searchQuery := helpers.ParseConditions(data.Conditions)
	log.Info().Msg(fmt.Sprintf("searchQuery: %#v", searchQuery))

	var cars []models.Car

	// Querying to find Cars.
	err := config.DB.Where(searchQuery).Find(&cars)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    cars,
	})
}
