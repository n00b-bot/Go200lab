package restaurantstorage

import (
	"context"
	restaurantsmodel "food/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantsmodel.RestaurantCreate) error {
	if err := s.db.Create(data); err != nil {
		return err.Error
	}

	return nil
}
