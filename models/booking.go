package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	RoomID    uint   `json:"room_id"`
	Room      Room   `gorm:"foreignKey:RoomID"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"` // "pending", "confirmed", "cancelled"
}
