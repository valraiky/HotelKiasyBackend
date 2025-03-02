package repositories

import (
	"golang-crud/config"
	"golang-crud/models"
)

type AvailabilityRepository interface {
	Create(availability *models.Availability) error
	FindAll() ([]models.Availability, error)
	FindByID(id uint) (*models.Availability, error)
	Update(availability *models.Availability) error
	Delete(id uint) error
	FindByRoomAndDates(roomID uint, startDate string, endDate string) ([]models.Availability, error)
}

type availabilityRepository struct{}

func NewAvailabilityRepository() AvailabilityRepository {
	return &availabilityRepository{}
}

func (r *availabilityRepository) Create(availability *models.Availability) error {
	return config.DB.Create(availability).Error
}

func (r *availabilityRepository) FindAll() ([]models.Availability, error) {
	var availabilities []models.Availability
	err := config.DB.Find(&availabilities).Error
	return availabilities, err
}

func (r *availabilityRepository) FindByID(id uint) (*models.Availability, error) {
	var availability models.Availability
	err := config.DB.First(&availability, id).Error
	return &availability, err
}

func (r *availabilityRepository) Update(availability *models.Availability) error {
	return config.DB.Save(availability).Error
}

func (r *availabilityRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Availability{}, id).Error
}

func (r *availabilityRepository) FindByRoomAndDates(roomID uint, startDate string, endDate string) ([]models.Availability, error) {
	var availabilities []models.Availability
	err := config.DB.Where("room_id = ? AND start_date <= ? AND end_date >= ?", roomID, endDate, startDate).Find(&availabilities).Error
	return availabilities, err
}
