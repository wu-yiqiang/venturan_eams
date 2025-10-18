package request

import (
	"venturan/global/serviceErrors"
)

type Register struct {
	Name       string `form:"name" json:"name" binding:"required"`
	NickName   string `form:"nick_name" json:"nick_name"`
	Email      string `form:"email" json:"email" binding:"required,email"`
	Avatar     string `form:"avatar" json:"avatar"`
	Mobile     string `form:"mobile" json:"mobile"`
	Password   string `form:"password" json:"password" binding:"required"`
	EmployTime string `form:"employ_time" json:"employ_time"`
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
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required":    serviceErrors.EmailAddrIsNotEmpty,
		"email.email":       serviceErrors.EmailFormaterError,
		"password.required": serviceErrors.UserPasswordIsNotEmpty,
	}
}
