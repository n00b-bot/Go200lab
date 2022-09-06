package restaurantlikestorage

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	rs := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}
	var listLike []sqlData
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Select("restaurant_id,count(restaurant_id) as count").Where("restaurant_id in (?)", ids).Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	for _, item := range listLike {
		rs[item.RestaurantId] = item.LikeCount
	}
	return rs, nil
}
