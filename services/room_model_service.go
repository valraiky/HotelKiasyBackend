package services

import (
	"golang-crud/models"
	"golang-crud/repositories"
)

type RoomModelService interface {
	CreateRoomModel(roomModel *models.RoomModel) error
	GetRoomModels() ([]models.RoomModel, error)
	GetRoomModelByID(id uint) (*models.RoomModel, error)
	UpdateRoomModel(roomModel *models.RoomModel) error
	DeleteRoomModel(id uint) error
}

type roomModelService struct {
	repo repositories.RoomModelRepository
}

func NewRoomModelService(repo repositories.RoomModelRepository) RoomModelService {
	return &roomModelService{repo: repo}
}

func (s *roomModelService) CreateRoomModel(roomModel *models.RoomModel) error {
	return s.repo.Create(roomModel)
}

func (s *roomModelService) GetRoomModels() ([]models.RoomModel, error) {
	return s.repo.FindAll()
}

func (s *roomModelService) GetRoomModelByID(id uint) (*models.RoomModel, error) {
	return s.repo.FindByID(id)
}

func (s *roomModelService) UpdateRoomModel(roomModel *models.RoomModel) error {
	return s.repo.Update(roomModel)
}

func (s *roomModelService) DeleteRoomModel(id uint) error {
	return s.repo.Delete(id)
}
