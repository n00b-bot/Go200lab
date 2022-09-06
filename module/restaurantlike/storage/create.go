package restaurantlikestorage

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
