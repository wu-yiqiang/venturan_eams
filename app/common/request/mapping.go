package request

import "venturan/global/serviceErrors"

type MappingCreate struct {
	Code            string `form:"code" json:"code" example:"" binding:"required"`
	Name            string `form:"name" json:"name"  example:"" binding:"required"`
	Value           uint   `form:"value" json:"value"  example:"1" binding:"required"`
	Color           string `json:"color" Gorm:"comment:文本颜色"`
	BackgroundColor string `json:"backgroundColor" Gorm:"comment:背景颜色"`
	Description     string `json:"description" gorm:"type:varchar(2500)"`
}

func (mappingCreate MappingCreate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"code.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":  serviceErrors.NameIsNotEmpty,
		"value.required": serviceErrors.ValueIsNotEmpty,
	}
}

type MappingUpdate struct {
	ID uint `json:"id" binding:"required"`
	MappingCreate
}

func (mappingUpdate MappingUpdate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":    serviceErrors.IdIsNotEmpty,
		"code.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":  serviceErrors.NameIsNotEmpty,
		"value.required": serviceErrors.ValueIsNotEmpty,
	}
}
