package controllers

import (
	"golang-crud/models"
	"golang-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AvailabilityController struct {
	service services.AvailabilityService
}

func NewAvailabilityController(service services.AvailabilityService) *AvailabilityController {
	return &AvailabilityController{service: service}
}

func (c *AvailabilityController) CreateAvailability(ctx *gin.Context) {
	var availability models.Availability
	ctx.BindJSON(&availability)
	if err := c.service.CreateAvailability(&availability); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, availability)
}

func (c *AvailabilityController) GetAvailabilities(ctx *gin.Context) {
	availabilities, _ := c.service.GetAvailabilities()
	ctx.JSON(http.StatusOK, availabilities)
}

func (c *AvailabilityController) GetAvailabilityByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	availability, err := c.service.GetAvailabilityByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Disponibilité non trouvée"})
		return
	}
	ctx.JSON(http.StatusOK, availability)
}

func (c *AvailabilityController) UpdateAvailability(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	var availability models.Availability
	ctx.BindJSON(&availability)
	availability.ID = uint(id)
	if err := c.service.UpdateAvailability(&availability); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, availability)
}

func (c *AvailabilityController) DeleteAvailability(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	if err := c.service.DeleteAvailability(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *AvailabilityController) GetRoomAvailability(ctx *gin.Context) {
	roomID, err := strconv.ParseUint(ctx.Query("room_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de chambre invalide"})
		return
	}
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	availabilities, err := c.service.GetRoomAvailability(uint(roomID), startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des disponibilités"})
		return
	}

	ctx.JSON(http.StatusOK, availabilities)
}
