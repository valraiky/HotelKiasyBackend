package routes

import (
	"golang-crud/controllers"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine, service services.BookingService) {
	controller := controllers.NewBookingController(service)
	bookingRoutes := router.Group("/bookings")
	{
		bookingRoutes.POST("", controller.CreateBooking)
		bookingRoutes.GET("", controller.GetBookings)
		bookingRoutes.GET("/:id", controller.GetBookingByID)
		bookingRoutes.PUT("/:id", controller.UpdateBooking)
		bookingRoutes.DELETE("/:id", controller.CancelBooking)
	}
}
