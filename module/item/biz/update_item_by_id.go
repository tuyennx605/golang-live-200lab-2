package biz

import (
	"context"
	"errors"
	"todo-list/common"
	"todo-list/module/item/model"
)

type UpdateItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemBiz struct {
	store     UpdateItemStore
	requester common.Requester
}

func NewUpdateItemBiz(store UpdateItemStore, requester common.Requester) *updateItemBiz {
	return &updateItemBiz{store, requester}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	deletedStatus := model.ItemStatusDeleted

	if data.Status == &deletedStatus {
		return model.ErrItemIsDeleted
	}

	isOwner := biz.requester.GetUserId() == data.UserId
	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission(errors.New("no permission"))
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
