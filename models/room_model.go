package models

import "gorm.io/gorm"

type RoomModel struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
