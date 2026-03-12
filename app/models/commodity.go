package models

type Commodity struct {
	ID
	Name
	Code
	Price     int64 `json:"price" Gorm:"price;comment:价格（分）"`
	Sales     uint  `json:"sales" Gorm:"sales;comment:销量"`
	Inventory uint  `json:"inventory" Gorm:"inventory;comment:库存数"`
	Status
	FileName
	Description
	Timestamps
	IsDeleted
}
