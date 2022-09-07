package restaurantlikestorage

import (
	"context"
	restaurantmodel "food/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) UpLike(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update(
		"liked_count", gorm.Expr("liked_count + ?", 1),
	).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DownLike(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update(
		"liked_count", gorm.Expr("liked_count - ?", 1),
	).Error; err != nil {
		return err
	}
	return nil
}
