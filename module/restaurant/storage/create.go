package restaurantstorage

import (
	"context"
	restaurantsmodel "food/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantsmodel.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
