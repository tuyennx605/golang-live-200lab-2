package model

import (
	"fmt"
	"todo-list/common"
	"todo-list/module/user/model"
)

const (
	EntityName = "UserLikeItem"
)

type Like struct {
	common.SQLModel
	UserId int               `json:"user_id" gorm:"column:user_id;"`
	ItemId int               `json:"item_id" gorm:"column:item_id;"`
	User   *model.SimpleUser `json:"-" gorm:"foreignKey:UserId;"` // foreignkey de ket noi
}

func (Like) TableName() string { return "user_like_items" }

func ErrCannotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this item"),
		fmt.Sprintf("ErrCannotLikeItem"),
	)
}

func ErrCannotUnLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this item"),
		fmt.Sprintf("ErrCannotUnLikeItem"),
	)
}

func ErrDidNotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Did not like this item"),
		fmt.Sprintf("ErrDidNotLikeItem"),
	)
}
