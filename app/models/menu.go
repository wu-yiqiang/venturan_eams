package models

type Menu struct {
	ID
	Name
	Path         string `json:"-" Gorm:"not null;comment:路径"`
	ParentMenuId uint   `json:"-" Gorm:"not null;comment:父级ID"`
	Code         string `json:"-" Gorm:"not null;comment:唯一编码"`
	MenuStatus   uint   `json:"-" Gorm:"not null;comment:菜单状态"`
	IconName     string `json:"-" Gorm:"default:'';comment:图标"`
	Description
	Timestamps
	IsDeleted
}
