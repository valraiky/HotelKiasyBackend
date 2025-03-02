package routes

import (
	"golang-crud/controllers"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func RoomModelRoutes(router *gin.Engine, service services.RoomModelService) {
	controller := controllers.NewRoomModelController(service)
	roomModelRoutes := router.Group("/room-models")
	{
		roomModelRoutes.POST("", controller.CreateRoomModel)
		roomModelRoutes.GET("", controller.GetRoomModels)
		roomModelRoutes.GET("/:id", controller.GetRoomModelByID)
		roomModelRoutes.PUT("/:id", controller.UpdateRoomModel)
		roomModelRoutes.DELETE("/:id", controller.DeleteRoomModel)
	}
}
