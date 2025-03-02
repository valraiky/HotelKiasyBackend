package controllers

import (
	"golang-crud/models"
	"golang-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	service services.RoomService
}

func NewRoomController(service services.RoomService) *RoomController {
	return &RoomController{service: service}
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
	var room models.Room
	ctx.BindJSON(&room)
	if err := c.service.CreateRoom(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, room)
}

func (c *RoomController) GetRooms(ctx *gin.Context) {
	rooms, _ := c.service.GetRooms()
	ctx.JSON(http.StatusOK, rooms)
}

func (c *RoomController) GetRoomByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	room, err := c.service.GetRoomByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Chambre non trouvée"})
		return
	}
	ctx.JSON(http.StatusOK, room)
}

func (c *RoomController) UpdateRoom(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	var room models.Room
	ctx.BindJSON(&room)
	room.ID = uint(id)
	if err := c.service.UpdateRoom(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, room)
}

func (c *RoomController) DeleteRoom(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	if err := c.service.DeleteRoom(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
