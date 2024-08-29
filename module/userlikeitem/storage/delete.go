package storage

import (
	"context"
	"todo-list/common"
	"todo-list/module/userlikeitem/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Delete(ctx context.Context, userId, itemId int) error {
	var data model.Like

	if err := s.db.Table(data.TableName()).Where("user_id = ? and item_id = ?", userId, itemId).Delete(nil).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}
	return nil
}
