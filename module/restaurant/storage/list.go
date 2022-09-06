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
			db = db.Where("user_id = ?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if paging.FakeCursor != "" {
		uid, err := common.FromBase58(paging.FakeCursor)
		if err != nil {
			return nil, err
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&rs).Error; err != nil {
		return nil, err
	}
	if len(rs) > 0 {
		last := rs[len(rs)-1]
		last.GenUID(common.RestaurantType)
		paging.NextCursor = last.FakeID.String()
	}
	return rs, nil
}
