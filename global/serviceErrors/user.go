package serviceErrors

import "venturan/app/common/response"

// User以1020000开头
var (
	UserNotFound                  = response.Response{1020000, nil, "用户不存在"}
	UserNotPermission             = response.Response{1020001, nil, "该用户没有权限"}
	UserIdNotEmpty                = response.Response{1020002, nil, "用户ID不能为空"}
	UserInfoNotExist              = response.Response{1020003, nil, "用户信息不存在"}
	UserIsLocked                  = response.Response{1020004, nil, "该账号已被锁定"}
	UserLogoutFail                = response.Response{1020005, nil, "用户注销失败"}
	UserIsExist                   = response.Response{1020006, nil, "用户已存在"}
	UserPasswordError             = response.Response{1020007, nil, "用户密码错误"}
	UserPasswordIsNotEmpty        = response.Response{1020008, nil, "用户密码不能为空"}
	UserIsNotExistOrPasswordError = response.Response{1020009, nil, "帐号不存在或密码错误"}
	UserCreateFailed              = response.Response{1020010, nil, "用户创建失败"}
	UserNameIsNotEmpty            = response.Response{1020011, nil, "用户名不能为空"}
)
