package request

import (
	"venturan/global/serviceErrors"
)

type Register struct {
	Name         string `form:"name" json:"name" example:"Sutter" binding:"required"`
	NickName     string `form:"nick_name" example:"Sutter" json:"nick_name"`
	Email        string `form:"email" json:"email" example:"sutter.wu@itforce-tech.com" binding:"required,email"`
	Password     string `form:"password" json:"password"example:"1234@Abcd" binding:"required"`
	Avatar       string `form:"avatar" example:"" json:"avatar"`
	Mobile       string `form:"mobile" example:"13557080956" json:"mobile"`
	DepartmentId uint   `form:"department_id" example:"1" json:"department_id"`
	PositionId   uint   `form:"position_id" example:"1" json:"position_id"`
	EmployDate   string `form:"employ_date" example:"2025-09-05" json:"employ_date"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    serviceErrors.EmailAddrIsNotEmpty,
		"email.email":       serviceErrors.EmailFormaterError,
		"password.required": serviceErrors.UserPasswordIsNotEmpty,
		"name.required":     serviceErrors.UserNameIsNotEmpty,
	}
}

type Login struct {
	Email    string `form:"email" json:"email" example:"sutter.wu@itforce-tech.com" binding:"required,email"`
	Password string `form:"password" json:"password" example:"1234@Abcd" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    serviceErrors.EmailAddrIsNotEmpty,
		"email.email":       serviceErrors.EmailFormaterError,
		"password.required": serviceErrors.UserPasswordIsNotEmpty,
	}
}
