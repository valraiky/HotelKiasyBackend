package services

import (
	"golang-crud/models"
	"golang-crud/repositories"
)

type BookingService interface {
	CreateBooking(booking *models.Booking) error
	GetBookings() ([]models.Booking, error)
	GetBookingByID(id uint) (*models.Booking, error)
	UpdateBooking(booking *models.Booking) error
	CancelBooking(id uint) error
}

type bookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) CreateBooking(booking *models.Booking) error {
	booking.Status = "pending"
	return s.repo.Create(booking)
}

func (s *bookingService) GetBookings() ([]models.Booking, error) {
	return s.repo.FindAll()
}

func (s *bookingService) GetBookingByID(id uint) (*models.Booking, error) {
	return s.repo.FindByID(id)
}

func (s *bookingService) UpdateBooking(booking *models.Booking) error {
	return s.repo.Update(booking)
}

func (s *bookingService) CancelBooking(id uint) error {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	booking.Status = "cancelled"
	return s.repo.Update(booking)
}
