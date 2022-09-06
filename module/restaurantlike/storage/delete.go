package restaurantlikestorage

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, restaurantId int) error {
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where(
		"user_id=? and restaurant_id=?", userId, restaurantId,
	).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
