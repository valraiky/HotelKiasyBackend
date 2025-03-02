package services

import (
	"golang-crud/models"
	"golang-crud/repositories"
)

type AvailabilityService interface {
	CreateAvailability(availability *models.Availability) error
	GetAvailabilities() ([]models.Availability, error)
	GetAvailabilityByID(id uint) (*models.Availability, error)
	UpdateAvailability(availability *models.Availability) error
	DeleteAvailability(id uint) error
	GetRoomAvailability(roomID uint, startDate string, endDate string) ([]models.Availability, error)
}

type availabilityService struct {
	repo repositories.AvailabilityRepository
}

func NewAvailabilityService(repo repositories.AvailabilityRepository) AvailabilityService {
	return &availabilityService{repo: repo}
}

func (s *availabilityService) CreateAvailability(availability *models.Availability) error {
	return s.repo.Create(availability)
}

func (s *availabilityService) GetAvailabilities() ([]models.Availability, error) {
	return s.repo.FindAll()
}

func (s *availabilityService) GetAvailabilityByID(id uint) (*models.Availability, error) {
	return s.repo.FindByID(id)
}

func (s *availabilityService) UpdateAvailability(availability *models.Availability) error {
	return s.repo.Update(availability)
}

func (s *availabilityService) DeleteAvailability(id uint) error {
	return s.repo.Delete(id)
}

func (s *availabilityService) GetRoomAvailability(roomID uint, startDate string, endDate string) ([]models.Availability, error) {
	return s.repo.FindByRoomAndDates(roomID, startDate, endDate)
}
