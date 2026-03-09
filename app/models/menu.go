package models

type Menu struct {
	ID
	Name
	Code
	Path         string `json:"path" Gorm:"not null;comment:路径"`
	ParentMenuId uint   `json:"parentMenuId" Gorm:"not null;comment:父级ID"`
	MenuStatus   uint   `json:"menuStatus" Gorm:"not null;comment:菜单状态"`
	IconName     string `json:"iconName" Gorm:"default:'';comment:图标"`
	Roles        []Role `gorm:"many2many:role_menu;"`
	Description
	Timestamps
	IsDeleted
}
