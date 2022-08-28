package restaurantstorage

import (
	"context"
	restaurantmodel "food/module/restaurant/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{
		"status": 0,
	}); err != nil {
		return err.Error
	}
	return nil

}