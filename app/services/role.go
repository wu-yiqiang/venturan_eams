package services

import (
	"venturan/app/common/request"
	"venturan/app/models"
	"venturan/global"
)

type roleService struct {
}

var RoleService = new(roleService)

func (roleService *roleService) RolePage(params request.CommonPageQueryForm) (err error, results *PaginationResult) {
	var roles []models.Role
	query := global.App.DB.Scopes(QueryIsNoeDeleted).Where("name LIKE ? OR code LIKE ? OR description LIKE ?", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%").Find(&roles)
	rolePageResults, err := Pagination(query, params.PageNo, params.PageSize, &roles)
	if err != nil {
		return
	}
	return nil, rolePageResults
}
