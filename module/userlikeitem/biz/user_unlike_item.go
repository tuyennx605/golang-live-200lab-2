package biz

import (
	"context"
	"log"
	"todo-list/common"
	"todo-list/module/userlikeitem/model"
)

type UserUnLikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type DecreaseLikeCount interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikeItemBiz struct {
	store     UserUnLikeItemStore
	itemStore DecreaseLikeCount
}

func NewUserUnLikeItemBiz(store UserUnLikeItemStore, itemStore DecreaseLikeCount) *userUnLikeItemBiz {
	return &userUnLikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userUnLikeItemBiz) UnLikeItem(ctx context.Context, userId, itemId int) error {

	_, err := biz.store.Find(ctx, userId, itemId)

	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnLikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnLikeItem(err)
	}

	//
	go func() {
		defer common.Recovery()
		if err := biz.itemStore.DecreaseLikeCount(ctx, itemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
