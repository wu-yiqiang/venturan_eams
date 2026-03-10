package models

type Mapping struct {
	ID
	Name
	Code
	Value           uint   `json:"value" Gorm:"not null;comment:字典值"`
	Color           string `json:"color" Gorm:"comment:文本颜色"`
	BackgroundColor string `json:"backgroundColor" Gorm:"comment:背景颜色"`
	Description
	Timestamps
	IsDeleted
}
