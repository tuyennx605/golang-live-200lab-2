package biz

import (
	"context"
	"todo-list/common"
	usermodel "todo-list/module/user/model"
	"todo-list/module/userlikeitem/model"
)

type ListUserLikeItemStore interface {
	ListUser(ctx context.Context, itemId int, paging *common.Paging) ([]usermodel.SimpleUser, error)
}

type listUserLikeItemBiz struct {
	store ListUserLikeItemStore
}

func NewListUserLikeItemBiz(store ListUserLikeItemStore) *listUserLikeItemBiz {
	return &listUserLikeItemBiz{store: store}
}

func (biz *listUserLikeItemBiz) ListUser(ctx context.Context, itemId int, paging *common.Paging) ([]usermodel.SimpleUser, error) {

	result, err := biz.store.ListUser(ctx, itemId, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return result, nil
}
