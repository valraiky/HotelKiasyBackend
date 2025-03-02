package controllers

import (
	"golang-crud/models"
	"golang-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service services.BookingService
}

func NewBookingController(service services.BookingService) *BookingController {
	return &BookingController{service: service}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var booking models.Booking
	ctx.BindJSON(&booking)
	if err := c.service.CreateBooking(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, booking)
}

func (c *BookingController) GetBookings(ctx *gin.Context) {
	bookings, _ := c.service.GetBookings()
	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) GetBookingByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	booking, err := c.service.GetBookingByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Réservation non trouvée"})
		return
	}
	ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) UpdateBooking(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	var booking models.Booking
	ctx.BindJSON(&booking)
	booking.ID = uint(id)
	if err := c.service.UpdateBooking(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) CancelBooking(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	if err := c.service.CancelBooking(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
