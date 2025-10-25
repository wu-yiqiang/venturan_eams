package models

import "strconv"

type User struct {
	ID
	Name
	NickName       string `json:"nick_name" gorm:"default:'';comment:用户别称"`
	Email          string `json:"email" gorm:"uniqueIndex;default:'';comment:邮箱地址"`
	Password       string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Avatar         string `json:"avatar" gorm:"default:'';comment:头像"`
	Mobile         string `json:"mobile" gorm:"default:'';index;comment:用户手机号"`
	EmployDate     string `json:"employ_date" gorm:"not null;comment:入职时间"`
	ResignDate     string `json:"resign_date" gorm:";comment:离职时间"`
	DepartmentId   uint   `json:"department_id"`
	DepartmentName string `gorm:"->" json:"department_name"`
	PositionId     uint   `json:"position_id"`
	PositionName   string `gorm:"->" json:"position_name"`
	Timestamps
	IsDeleted
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
