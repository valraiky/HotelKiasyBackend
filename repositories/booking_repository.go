package repositories

import (
	"golang-crud/config"
	"golang-crud/models"
)

type BookingRepository interface {
	Create(booking *models.Booking) error
	FindAll() ([]models.Booking, error)
	FindByID(id uint) (*models.Booking, error)
	Update(booking *models.Booking) error
	Delete(id uint) error
}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}

func (r *bookingRepository) Create(booking *models.Booking) error {
	return config.DB.Create(booking).Error
}

func (r *bookingRepository) FindAll() ([]models.Booking, error) {
	var bookings []models.Booking
	err := config.DB.Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) FindByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	err := config.DB.First(&booking, id).Error
	return &booking, err
}

func (r *bookingRepository) Update(booking *models.Booking) error {
	return config.DB.Save(booking).Error
}

func (r *bookingRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Booking{}, id).Error
}
