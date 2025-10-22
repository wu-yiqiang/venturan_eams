package serviceErrors

import "venturan/app/common/response"

// 公共部分以1010001开头
var (
	RequestParamsError       = response.Response{1010000, nil, "请求参数错误"}
	IdIsNotExist             = response.Response{1010001, nil, "ID不能为空"}
	TokenIsNotEmpty          = response.Response{1010002, nil, "token不为空"}
	TokenIsNotExist          = response.Response{1010003, nil, "token不存在"}
	TokenIsExpired           = response.Response{1010004, nil, "token过期"}
	EmailFormaterError       = response.Response{1010005, nil, "邮箱格式错误"}
	EmailAddrIsNotEmpty      = response.Response{1010006, nil, "邮箱地址不能为空"}
	TokentokenIssuanceFailed = response.Response{1010007, nil, "token签发失败"}
	EmailAddIsExisted        = response.Response{1010008, nil, "邮箱地址已存在"}
	FrequentRequests         = response.Response{1010009, nil, "接口请求频繁，请稍后重试"}
)
