package restaurantlikestorage

import (
	"context"
	"fmt"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:05:06.999999"

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

func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context, conditions map[string]interface{}, filter *restaurantlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error) {
	var rs []restaurantlikemodel.Like
	db := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	db = db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		time, err := time.Parse(timeLayout, string(base58.Decode(paging.FakeCursor)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at < ?", time.Format("2006-01-02 15:05:01"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&rs).Error; err != nil {
		return nil, err
	}
	users := make([]common.SimpleUser, len(rs))

	for i, u := range rs {
		rs[i].User.CreatedAt = u.CreatedAt
		users[i] = *u.User
		if i == len(rs)-1 {
			cur := base58.Encode([]byte(fmt.Sprint(u.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cur
		}
	}
	return users, nil
}
