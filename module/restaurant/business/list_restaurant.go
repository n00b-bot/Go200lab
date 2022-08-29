package restaurantsbusiness

import (
	"context"
	"food/common"
	restaurantmodel "food/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context.Context, *restaurantmodel.Filter, *common.Paging, ...string) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBus struct {
	store ListRestaurantStore
}

func NewListRestaurantBus(store ListRestaurantStore) *ListRestaurantBus {
	return &ListRestaurantBus{
		store: store,
	}
}

func (l *ListRestaurantBus) List(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	rs, err := l.store.ListDataWithCondition(context, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.Entity, err)
	}
	return rs, nil
}
