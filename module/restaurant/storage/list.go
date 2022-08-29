package restaurantstorage

import (
	"context"
	"food/common"
	restaurantmodel "food/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var rs []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	offset := (paging.Page - 1) * paging.Limit
	if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}
