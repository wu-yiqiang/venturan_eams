package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Email    string `json:"email" gorm:"comment:邮箱地址"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	IsDeleted
}
