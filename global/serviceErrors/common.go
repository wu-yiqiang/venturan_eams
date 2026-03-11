package serviceErrors

import "venturan/app/common/response"

// 公共部分以1010001开头
var (
	RequestParamsError  = response.Response{1010000, nil, "请求参数错误"}
	IdIsNotExist        = response.Response{1010001, nil, "ID不存在"}
	TokenIsExpired      = response.Response{1010002, nil, "token过期或者不存在"}
	EmailFormaterError  = response.Response{1010003, nil, "邮箱格式错误"}
	EmailAddrIsNotEmpty = response.Response{1010004, nil, "邮箱地址不能为空"}
	TokenIssuanceFailed = response.Response{1010005, nil, "token签发失败"}
	EmailAddIsExisted   = response.Response{1010006, nil, "邮箱地址已存在"}
	FrequentRequests    = response.Response{1010007, nil, "接口请求频繁，请稍后重试"}
	DeleteError         = response.Response{1010008, nil, "删除失败"}
	QueryError          = response.Response{1010009, nil, "查询列表失败"}
	CreateError         = response.Response{1010010, nil, "创建失败"}
	UpdateError         = response.Response{1010011, nil, "更新失败"}
	PageSizeIsNotEmpty  = response.Response{1010012, nil, "pageSize不能为空"}
	PageNoIsNotEmpty    = response.Response{1010013, nil, "pageNo不能为空"}
	CodeIsNotEmpty      = response.Response{1010014, nil, "编码不能为空"}
	NameIsNotEmpty      = response.Response{1010015, nil, "名字不能为空"}
	ValueIsNotEmpty     = response.Response{1010016, nil, "值不能为空"}
	QueryDetailsError   = response.Response{1010017, nil, "查询详情失败"}
	IdIsNotEmpty        = response.Response{1010018, nil, "ID不能为空"}
)
