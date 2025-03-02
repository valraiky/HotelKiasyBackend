package main

import (
	"golang-crud/config"
	"golang-crud/repositories"
	"golang-crud/routes"
	"golang-crud/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connexion à la base de données
	config.ConnectDatabase()

	// Initialisation des composants
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	roomModelRepo := repositories.NewRoomModelRepository()
	roomModelService := services.NewRoomModelService(roomModelRepo)

	roomRepo := repositories.NewRoomRepository()
	roomService := services.NewRoomService(roomRepo)

	bookingRepo := repositories.NewBookingRepository()
	bookingService := services.NewBookingService(bookingRepo)

	availabilityRepo := repositories.NewAvailabilityRepository()
	availabilityService := services.NewAvailabilityService(availabilityRepo)

	// Configuration des routes
	router := gin.Default()
	routes.UserRoutes(router, userService)
	routes.RoomModelRoutes(router, roomModelService)
	routes.RoomRoutes(router, roomService)
	routes.BookingRoutes(router, bookingService)
	routes.AvailabilityRoutes(router, availabilityService)

	// Démarrage du serveur
	router.Run(":8080")
}
