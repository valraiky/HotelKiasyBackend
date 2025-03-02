package services

import (
	"errors"
	"golang-crud/models"
	"golang-crud/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUser(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("nom et email obligatoires")
	}
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) UpdateUser(user *models.User) error { // Implémentation de la méthode UpdateUser
	if user.Name == "" || user.Email == "" {
		return errors.New("nom et email obligatoires")
	}
	return s.repo.Update(user)
}
