package biz

import (
	"context"
	"log"
	"todo-list/common"
	"todo-list/module/userlikeitem/model"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseLikeCount interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeItemBiz struct {
	store     UserLikeItemStore
	itemStore IncreaseLikeCount
}

func NewUserLikeItemBiz(store UserLikeItemStore, itemStore IncreaseLikeCount) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	//
	go func() {
		defer common.Recovery()
		if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
			log.Println(err)
		}

	}()

	return nil
}
