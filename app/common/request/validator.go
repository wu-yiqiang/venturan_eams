package request

import (
	"github.com/go-playground/validator/v10"
	"venturan/app/common/response"
	"venturan/global/serviceErrors"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]response.Response

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) response.Response {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			// 若 request 结构体实现 Validator 接口即可实现自定义错误信息
			if isValidator {
				if errorType, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return errorType
				}
			}
			return serviceErrors.RequestParamsError
		}
	}
	return serviceErrors.RequestParamsError
}
