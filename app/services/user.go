package services

import (
	"errors"
	"strconv"
	"time"
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
	user = models.User{Email: params.Email, Name: models.Name{params.Name}, DepartmentId: params.DepartmentId, PositionId: params.DepartmentId, Mobile: params.Mobile, Avatar: params.Avatar, EmployDate: params.EmployDate, NickName: params.NickName, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	if err != nil {
		err = errors.New(serviceErrors.UserCreateFailed.Msg)
	}
	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("email = ?", params.Email).Select("users.*, departments.name as department_name, positions.name as position_name").
		Joins("LEFT JOIN departments ON users.department_id = departments.id").
		Joins("LEFT JOIN positions ON users.position_id = positions.id").First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New(serviceErrors.UserIsNotExistOrPasswordError.Msg)
	}
	return
}

func (userService *userService) Delete(userId uint) (err error, user *models.User) {
	err = global.App.DB.Where("id = ?", userId).First(&user).Update("is_deleted", 1).Error
	if err != nil {
		err = errors.New(serviceErrors.DeleteError.Msg)
	}
	return
}

func (userService *userService) DeleteResgin() (err error) {
	users := []*models.User{}
	err = global.App.DB.Where("is_deleted = 0 AND resign_date < ?", time.Now()).Find(&users).Update("is_deleted", 1).Error
	if err != nil {
		return
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
