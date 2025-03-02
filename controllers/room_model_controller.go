package controllers

import (
	"golang-crud/models"
	"golang-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomModelController struct {
	service services.RoomModelService
}

func NewRoomModelController(service services.RoomModelService) *RoomModelController {
	return &RoomModelController{service: service}
}

func (c *RoomModelController) CreateRoomModel(ctx *gin.Context) {
	var roomModel models.RoomModel
	ctx.BindJSON(&roomModel)
	if err := c.service.CreateRoomModel(&roomModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, roomModel)
}

func (c *RoomModelController) GetRoomModels(ctx *gin.Context) {
	roomModels, _ := c.service.GetRoomModels()
	ctx.JSON(http.StatusOK, roomModels)
}

func (c *RoomModelController) GetRoomModelByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	roomModel, err := c.service.GetRoomModelByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Modèle de chambre non trouvé"})
		return
	}
	ctx.JSON(http.StatusOK, roomModel)
}

func (c *RoomModelController) UpdateRoomModel(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	var roomModel models.RoomModel
	ctx.BindJSON(&roomModel)
	roomModel.ID = uint(id)
	if err := c.service.UpdateRoomModel(&roomModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, roomModel)
}

func (c *RoomModelController) DeleteRoomModel(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	if err := c.service.DeleteRoomModel(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
