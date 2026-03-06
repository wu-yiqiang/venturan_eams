package services

import (
	"venturan/app/common/request"
	"venturan/app/models"
)

type menuService struct {
}

var MenuService = new(menuService)

func (menuService *menuService) MenuPage(params request.MenuPage) (err error, menus []models.Menu) {
	//var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	//if result.RowsAffected != 0 {
	//	err = errors.New(serviceErrors.EmailAddIsExisted.Msg)
	//	return
	//}
	//user = models.Menu{Email: params.Email, Name: models.Name{params.Name}, DepartmentId: params.DepartmentId, PositionId: params.DepartmentId, Mobile: params.Mobile, Avatar: params.Avatar, EmployDate: params.EmployDate, NickName: params.NickName, Password: utils.BcryptMake([]byte(params.Password))}
	//err = global.App.DB.Create(&user).Error
	//if err != nil {
	//	err = errors.New(serviceErrors.UserCreateFailed.Msg)
	//}
	return
}
