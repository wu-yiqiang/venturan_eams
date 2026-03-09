package models

type Button struct {
	ID
	Name
	Code
	Roles []Role `gorm:"many2many:role_button;"`
	Description
	Timestamps
	IsDeleted
}
