package request

import (
	"venturan/global/serviceErrors"
)

type ID struct {
	ID uint `form:"id" json:"id" example:"1" binding:"required"`
}

type CommonPageQueryForm struct {
	Search   string `form:"search" json:"search" example:""`
	PageSize int    `form:"pageSize" json:"pageSize" example:"10" binding:"required"`
	PageNo   int    `form:"pageNo" json:"pageNo" example:"1" binding:"required"`
}

func (commonPageQueryForm CommonPageQueryForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"pageNo.required":   serviceErrors.PageNoIsNotEmpty,
		"pageSize.required": serviceErrors.PageSizeIsNotEmpty,
	}
}
