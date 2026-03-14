package request

import "venturan/global/serviceErrors"

type MappingCreate struct {
	MappingType     string `form:"mappingType" json:"mappingType" example:"" binding:"required"`
	Name            string `form:"name" json:"name"  example:"" binding:"required"`
	MappingValue    uint   `form:"mappingValue" json:"mappingValue"  example:"1" binding:"required"`
	Color           string `json:"color" Gorm:"comment:文本颜色"`
	BackgroundColor string `json:"backgroundColor" Gorm:"comment:背景颜色"`
	Description     string `json:"description" gorm:"type:varchar(2500)"`
}

func (mappingCreate MappingCreate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mappingType.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":         serviceErrors.NameIsNotEmpty,
		"mappingValue.required": serviceErrors.ValueIsNotEmpty,
	}
}

type MappingUpdate struct {
	ID uint `json:"id" binding:"required"`
	MappingCreate
}

func (mappingUpdate MappingUpdate) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":           serviceErrors.IdIsNotEmpty,
		"mappingType.required":  serviceErrors.CodeIsNotEmpty,
		"name.required":         serviceErrors.NameIsNotEmpty,
		"mappingValue.required": serviceErrors.ValueIsNotEmpty,
	}
}

type MappingType struct {
	MappingType string `form:"mappingType" json:"mappingType" example:"" binding:"required"`
}

func (mappingType MappingType) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mappingType.required": serviceErrors.MappingTypeIsNotEmpty,
	}
}
