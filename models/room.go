package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomModelID uint      `json:"room_model_id"`
	RoomModel   RoomModel `gorm:"foreignKey:RoomModelID"`
	RoomNumber  string    `json:"room_number"`
	Available   bool      `json:"available"`
}
