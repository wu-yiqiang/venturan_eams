package models

type PlateNumber struct {
	ID
	Name
	UserId uint `json:"user_id"`
	Timestamps
	IsDeleted
}
