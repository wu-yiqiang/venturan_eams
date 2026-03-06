package request

import "venturan/global/serviceErrors"

type MenuPage struct {
	Keywords string `form:"keywords" json:"keywords" example:"" binding:"keywords"`
	PageSize uint   `form:"pageSize" json:"pageSize"  example:"10" binding:"required"`
	PageNo   uint   `form:"pageNo" json:"pageNo"  example:"1" binding:"required"`
}

func (menuPage MenuPage) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"pageSize.required": serviceErrors.PageSizeIsNotEmpty,
		"pageNo.required":   serviceErrors.PageNoIsNotEmpty,
	}
}
