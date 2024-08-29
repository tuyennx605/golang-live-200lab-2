package biz

import (
	"context"
	"todo-list/common"
	"todo-list/module/item/model"
)

type DeleteItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStore
}

func NewDeleteItemBiz(store DeleteItemStore) *deleteItemBiz {
	return &deleteItemBiz{store}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	deletedStatus := model.ItemStatusDeleted

	if data.Status == &deletedStatus {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}
	return nil
}
