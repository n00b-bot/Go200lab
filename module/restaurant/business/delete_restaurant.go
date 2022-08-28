package restaurantsbusiness

import (
	"context"
	"errors"
	restaurantmodel "food/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(context.Context, int) error
	FindDataWithCondition(context.Context, map[string]interface{}, ...string) (*restaurantmodel.Restaurant, error)
}

type DeleteRestaurantBus struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBus(store DeleteRestaurantStore) *DeleteRestaurantBus {
	return &DeleteRestaurantBus{
		store: store,
	}
}

func (d *DeleteRestaurantBus) Delete(context context.Context, id int) error {
	oldData, err := d.store.FindDataWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := d.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
