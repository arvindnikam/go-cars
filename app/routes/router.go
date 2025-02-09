package routes

import (
	car_controllers "cars/app/controllers/car"
	car_variant_controllers "cars/app/controllers/car_variant"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()
	route.SetTrustedProxies(nil)

	v1 := route.Group("/api/v1")
	{
		v1.POST("/cars/create", car_controllers.CreateCar)
		v1.GET("/cars/:car_id", car_controllers.ShowCar)
		v1.POST("/cars/search", car_controllers.SearchCars)
		v1.PUT("/cars/:car_id/update", car_controllers.UpdateCar)
		v1.DELETE("/cars/:car_id/delete", car_controllers.DeleteCar)

		car := v1.Group("/cars/:car_id")
		{
			car.POST("/car_variants/create", car_variant_controllers.CreateCarVariant)
			car.GET("/car_variants/:car_variant_id", car_variant_controllers.ShowCarVariant)
			car.POST("/car_variants/search", car_variant_controllers.SearchCarVariants)
			car.PUT("/car_variants/:car_variant_id/update", car_variant_controllers.UpdateCarVariant)
			car.DELETE("/car_variants/:car_variant_id/delete", car_variant_controllers.DeleteCarVariant)
		}
	}

	route.Run(":8880")
}
