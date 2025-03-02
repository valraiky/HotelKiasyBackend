package routes

import (
	"golang-crud/controllers"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func AvailabilityRoutes(router *gin.Engine, service services.AvailabilityService) {
	controller := controllers.NewAvailabilityController(service)
	availabilityRoutes := router.Group("/availabilities")
	{
		availabilityRoutes.POST("", controller.CreateAvailability)
		availabilityRoutes.GET("", controller.GetAvailabilities)
		availabilityRoutes.GET("/:id", controller.GetAvailabilityByID)
		availabilityRoutes.PUT("/:id", controller.UpdateAvailability)
		availabilityRoutes.DELETE("/:id", controller.DeleteAvailability)
		availabilityRoutes.GET("/room", controller.GetRoomAvailability) // Route pour récupérer les disponibilités d'une chambre
	}
}
