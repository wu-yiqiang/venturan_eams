package models

type Role struct {
	ID
	Name
	Code
	Menus      []Menu      `gorm:"many2many:role_menu;" json:"menus"`
	Connectors []Connector `gorm:"many2many:role_connector;" json:"connectors"`
	Buttons    []Button    `gorm:"many2many:role_button;" json:"buttons"`
	Users      []User      `gorm:"many2many:user_role;" `
	Description
	Timestamps
	IsDeleted
}
