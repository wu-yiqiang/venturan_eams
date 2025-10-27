package models

type CookBook struct {
	ID
	Name
	Price float32 `json:"price" gorm:"type:decimal(10,2)"`
	Image
	Description
	Timestamps
	IsDeleted
}
