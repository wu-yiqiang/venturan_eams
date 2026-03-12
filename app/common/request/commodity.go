package request

import "venturan/global/serviceErrors"

type CommodityCreate struct {
	Code        string `form:"code" json:"code" example:"" binding:"required"`
	Name        string `form:"name" json:"name"  example:"" binding:"required"`
	Price       int64  `json:"price" Gorm:"price;comment:价格（分）"`
	Sales       uint   `json:"sales" Gorm:"sales;comment:销量"`
	Inventory   uint   `json:"inventory" Gorm:"inventory;comment:库存数"`
	Status      int8   `json:"status" gorm:"index"`
	FileName    string `json:"fileName" Gorm:"not null;comment:文件名"`
	Description string `json:"description" gorm:"type:varchar(2500)"`
}

func (commodityCreate CommodityCreate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"code.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":  serviceErrors.NameIsNotEmpty,
		"value.required": serviceErrors.ValueIsNotEmpty,
	}
}

type CommodityUpdate struct {
	ID uint `json:"id" binding:"required"`
	CommodityCreate
}

func (commodityUpdate CommodityUpdate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":    serviceErrors.IdIsNotEmpty,
		"code.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":  serviceErrors.NameIsNotEmpty,
		"value.required": serviceErrors.ValueIsNotEmpty,
	}
}
