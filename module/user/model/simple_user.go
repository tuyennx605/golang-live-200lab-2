package model

import "todo-list/common"

type SimpleUser struct {
	common.SQLModel
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"role" gorm:"column:role;"`
	Status    string `json:"status" gorm:"column:status;"`
}

func (SimpleUser) TableName() string {
	return User{}.TableName()
}
