package models

import "gorm.io/gorm"

type Availability struct {
	gorm.Model
	RoomID    uint   `json:"room_id"`
	Room      Room   `gorm:"foreignKey:RoomID"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Available bool   `json:"available"`
}
