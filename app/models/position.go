package models

type Position struct {
	ID
	PositionName string `json:"position_name" gorm:"not null;comment:岗位名称"`
	Timestamps
	IsDeleted
}
