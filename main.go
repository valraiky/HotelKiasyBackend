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

	// Initialisation du routeur Gin
	router := gin.Default()

	// Définition des dépendances (Injection de dépendances)
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	// Définition des routes
	routes.UserRoutes(router, userService)

	// Lancement du serveur sur le port 8080
	router.Run(":8080")
}
