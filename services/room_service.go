package services

import (
	"golang-crud/models"
	"golang-crud/repositories"
)

type RoomService interface {
	CreateRoom(room *models.Room) error
	GetRooms() ([]models.Room, error)
	GetRoomByID(id uint) (*models.Room, error)
	UpdateRoom(room *models.Room) error
	DeleteRoom(id uint) error
}

type roomService struct {
	repo repositories.RoomRepository
}

func NewRoomService(repo repositories.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

func (s *roomService) CreateRoom(room *models.Room) error {
	return s.repo.Create(room)
}

func (s *roomService) GetRooms() ([]models.Room, error) {
	return s.repo.FindAll()
}

func (s *roomService) GetRoomByID(id uint) (*models.Room, error) {
	return s.repo.FindByID(id)
}

func (s *roomService) UpdateRoom(room *models.Room) error {
	return s.repo.Update(room)
}

func (s *roomService) DeleteRoom(id uint) error {
	return s.repo.Delete(id)
}
