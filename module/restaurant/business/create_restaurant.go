package restaurantsbusiness

import (
	"context"
	restaurantsmodel "food/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantsmodel.RestaurantCreate) error
}

type createRestaurantBus struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBus(store CreateRestaurantStore) createRestaurantBus {
	return createRestaurantBus{
		store: store,
	}

}

func (b *createRestaurantBus) CreateRestaurant(context context.Context, data *restaurantsmodel.RestaurantCreate) error {
	if err := b.store.Create(context, data); err != nil {
		return err
	}
	return nil
}
