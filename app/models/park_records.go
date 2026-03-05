package models

import "time"

type ParkRecord struct {
	ID
	PlateNumberId uint      `json:"plate_number_id"`
	InTime        time.Time `json:"in_time" Gorm:"column:in_time;default:now()"`
	OutTime       time.Time `json:"out_time" Gorm:"column:out_time;default:null"`
	Timestamps
	IsDeleted
}
