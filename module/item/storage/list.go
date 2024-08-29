package storage

import (
	"context"
	"todo-list/common"
	"todo-list/module/item/model"
)

func (s *sqlStore) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	requester := ctx.Value(common.CurrentUser).(common.Requester) // get value từ thằng ctx và ép kiểu thành common.Requester
	db = db.Where("user_id = ?", requester.GetUserId())

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Select("id").Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// nen lam duoi cau lenh count
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Select("*").Order("id desc").Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
