package routes

import (
	"golang-crud/controllers"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func RoomRoutes(router *gin.Engine, service services.RoomService) {
	controller := controllers.NewRoomController(service)
	roomRoutes := router.Group("/rooms")
	{
		roomRoutes.POST("", controller.CreateRoom)
		roomRoutes.GET("", controller.GetRooms)
		roomRoutes.GET("/:id", controller.GetRoomByID)
		roomRoutes.PUT("/:id", controller.UpdateRoom)
		roomRoutes.DELETE("/:id", controller.DeleteRoom)
	}
}
