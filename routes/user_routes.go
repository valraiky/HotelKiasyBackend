package routes

import (
	"golang-crud/controllers"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, service services.UserService) {
	controller := controllers.NewUserController(service)
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", controller.CreateUser)
		userRoutes.GET("", controller.GetUsers)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.DELETE("/:id", controller.DeleteUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
	}
}
