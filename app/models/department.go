package models

type Department struct {
	ID
	DepartmentName string `json:"department_name" gorm:"not null;comment:部门名称"`
	Users          []User `gorm:"many2many:department_user;"`
	Timestamps
	IsDeleted
}
