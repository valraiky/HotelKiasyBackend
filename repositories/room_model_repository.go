package repositories

import (
	"golang-crud/config"
	"golang-crud/models"
)

type RoomModelRepository interface {
	Create(roomModel *models.RoomModel) error
	FindAll() ([]models.RoomModel, error)
	FindByID(id uint) (*models.RoomModel, error)
	Update(roomModel *models.RoomModel) error
	Delete(id uint) error
}

type roomModelRepository struct{}

func NewRoomModelRepository() RoomModelRepository {
	return &roomModelRepository{}
}

func (r *roomModelRepository) Create(roomModel *models.RoomModel) error {
	return config.DB.Create(roomModel).Error
}

func (r *roomModelRepository) FindAll() ([]models.RoomModel, error) {
	var roomModels []models.RoomModel
	err := config.DB.Find(&roomModels).Error
	return roomModels, err
}

func (r *roomModelRepository) FindByID(id uint) (*models.RoomModel, error) {
	var roomModel models.RoomModel
	err := config.DB.First(&roomModel, id).Error
	return &roomModel, err
}

func (r *roomModelRepository) Update(roomModel *models.RoomModel) error {
	return config.DB.Save(roomModel).Error
}

func (r *roomModelRepository) Delete(id uint) error {
	return config.DB.Delete(&models.RoomModel{}, id).Error
}
