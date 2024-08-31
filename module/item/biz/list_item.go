package biz

import (
	"context"
	"todo-list/common"
	"todo-list/module/item/model"
)

type ListItemRepo interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	repo      ListItemRepo
	requester common.Requester
}

func NewListItemRepo(repo ListItemRepo, requester common.Requester) *listItemBiz {
	return &listItemBiz{repo, requester}
}

func (biz *listItemBiz) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester) // gán requester vào contex
	data, err := biz.repo.ListItem(ctxStore, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return data, nil
}
