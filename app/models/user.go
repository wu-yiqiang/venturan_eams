package models

import "strconv"

type User struct {
	ID
	Name       string `json:"name" gorm:"not null;comment:用户名称"`
	NickName   string `json:"nick_name" gorm:"default:'';comment:用户别称"`
	Email      string `json:"email" gorm:"default:'';comment:邮箱地址"`
	Avatar     string `json:"avatar" gorm:"default:'';comment:头像"`
	Mobile     string `json:"mobile" gorm:"default:'';index;comment:用户手机号"`
	Password   string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	EmployTime string `json:"employ_time" gorm:"not null;index;comment:入职时间"`
	Timestamps
	IsDeleted
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
