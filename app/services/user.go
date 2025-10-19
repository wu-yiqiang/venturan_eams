package services

import (
	"errors"
	"strconv"
	"venturan/app/common/request"
	"venturan/app/models"
	"venturan/global"
	"venturan/global/serviceErrors"
	"venturan/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New(serviceErrors.EmailAddIsExisted.Msg)
		return
	}
	user = models.User{Email: params.Email, Name: models.Name{params.Name}, DepartmentId: params.DepartmentId, PositionId: params.DepartmentId, Mobile: params.Mobile, Avatar: params.Avatar, EmployDay: params.EmployDay, NickName: params.NickName, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	if err != nil {
		err = errors.New(serviceErrors.UserCreateFailed.Msg)
	}
	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New(serviceErrors.UserIsNotExistOrPasswordError.Msg)
	}
	return
}

func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
