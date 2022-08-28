package restaurantstorage

import (
	"context"
	restaurantmodel "food/module/restaurant/model"
)

func (s *sqlStore) FindRestaurantWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data); err != nil {
		return nil, err.Error
	}
	return &data, nil
}
