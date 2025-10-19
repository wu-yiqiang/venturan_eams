package models

import "strconv"

type User struct {
	ID
	Name
	NickName     string `json:"nick_name" gorm:"default:'';comment:用户别称"`
	Email        string `json:"email" gorm:"default:'';comment:邮箱地址"`
	Avatar       string `json:"avatar" gorm:"default:'';comment:头像"`
	Mobile       string `json:"mobile" gorm:"default:'';index;comment:用户手机号"`
	Password     string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	EmployDay    string `json:"employ_day" gorm:"not null;index;comment:入职时间"`
	DepartmentId uint   `json:"department_id"`
	PositionId   uint   `json:"position_id"`
	Timestamps
	IsDeleted
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
