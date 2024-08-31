package storage

import (
	"context"
	"time"
	"todo-list/common"
	usermodel "todo-list/module/user/model"
	"todo-list/module/userlikeitem/model"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) ListUser(ctx context.Context, itemId int, paging *common.Paging) ([]usermodel.SimpleUser, error) {
	var result []model.Like

	db := s.db.Where("item_id = ?", itemId)
	if err := db.Table(model.Like{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// Seek Paging
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05.999999"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").Order("created_at desc").Limit(paging.Limit).Preload("User").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]usermodel.SimpleUser, len(result))
	for i := range users {
		users[i] = *result[i].User
	}

	if len(result) > 0 {
		paging.NextCursor = base58.Encode([]byte(result[len(result)-1].CreatedAt.Format(timeLayout)))
	}

	return users, nil
}

func (s *sqlStore) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) { // trả về map với itemId : bao nhiêu like
	result := make(map[int]int)
	type sqlData struct {
		ItemId int `gorm:"column:item_id"`
		Count  int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(model.Like{}.TableName()).Select("item_id, COUNT(item_id) as count").Where("item_id in (?)", ids).Group("item_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.ItemId] = item.Count
	}

	return result, nil
}
