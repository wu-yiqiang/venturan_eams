package models

import "time"

type ParkRecord struct {
	ID
	PlateNumberId uint      `json:"plate_number_id"`
	InTime        time.Time `json:"in_time" gorm:"column:in_time;default:now()"`
	OutTime       time.Time `json:"out_time" gorm:"column:out_time;default:null"`
	Timestamps
	IsDeleted
}
