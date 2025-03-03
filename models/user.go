package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	ImagePath string `json:"image_path"`
	Password  string `json:"-"` // Ajout de json:"-"
	Role      string `json:"role" gorm:"default:user"`
}
