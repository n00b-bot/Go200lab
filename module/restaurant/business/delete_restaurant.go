package restaurantsbusiness

import "context"

type DeleteRestaurantStore interface {
	Delete(context.Context, int) error
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
	if err := d.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
