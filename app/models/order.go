package models

type Order struct {
	ID
	Name
	SerialNumber
	Users []User `gorm:"many2many:user_role;" `
	Description
	Timestamps
	IsDeleted
}
