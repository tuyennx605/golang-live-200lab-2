package storage

import (
	"context"
	"todo-list/common"
	usermodel "todo-list/module/user/model"
	"todo-list/module/userlikeitem/model"
)

func (s *sqlStore) ListUser(ctx context.Context, itemId int, paging *common.Paging) ([]usermodel.SimpleUser, error) {
	var result []model.Like

	db := s.db.Where("item_id = ?", itemId)
	if err := db.Table(model.Like{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Select("*").Order("created_at desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Preload("User").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]usermodel.SimpleUser, len(result))
	for i := range users {
		users[i] = *result[i].User
	}
	return users, nil
}
