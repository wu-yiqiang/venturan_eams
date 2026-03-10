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
	err = global.App.DB.Where("email = ?", params.Email).Preload("Roles").Preload("Roles.Menus").Preload("Roles.Buttons").First(&user).Error
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

func (userService *userService) UserLogout(key string) (err error) {
	err = utils.ClearValue(key)
	if err != nil {
		err = errors.New(serviceErrors.UserLogoutFail.Msg)
	}
	return
}

func (userService *userService) FindUserByPlateNumber(plateNumber string) (err error, user models.User) {
	err = global.App.DB.Where("plate_number = ? And is_deleted = ?", plateNumber, 0).First(&user).Error
	if err != nil {
		return
	}
	return nil, user
}

func (userService *userService) UserPage(params request.UserPageQueryForm) (err error, results *PaginationResult) {
	var users []models.User
	query := global.App.DB.Scopes(QueryIsNoeDeleted).Where("name LIKE ? OR nick_name LIKE ? OR email LIKE ?", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%").Find(&users)
	userPageResults, err := Pagination(query, params.PageNo, params.PageSize, &users)
	if err != nil {
		return
	}
	return nil, userPageResults
}
