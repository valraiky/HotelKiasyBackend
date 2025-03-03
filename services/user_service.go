package services

import (
	"errors"
	"golang-crud/models"
	"golang-crud/repositories"
	"golang-crud/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	Login(email string, password string) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("nom, email et mot de passe obligatoires")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	if user.Role == "" {
		user.Role = "user"
	}

	return s.repo.Create(user)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("mot de passe incorrect")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
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
