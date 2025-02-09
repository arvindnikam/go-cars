package main

import (
	"cars/app/config"
	"cars/app/routes"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(DB)
	routes.Routes()
}
