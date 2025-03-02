package repositories

import (
	"golang-crud/config"
	"golang-crud/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Delete(id uint) error
	Update(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func (r *userRepository) Update(user *models.User) error { // Implémentation de la méthode Update
	return config.DB.Save(user).Error
}
