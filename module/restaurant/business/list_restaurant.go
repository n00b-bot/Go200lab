package restaurantsbusiness

import (
	"context"
	"food/common"
	restaurantmodel "food/module/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context.Context, *restaurantmodel.Filter, *common.Paging, ...string) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type ListRestaurantBus struct {
	store ListRestaurantStore
	like  LikeRestaurantStore
}

func NewListRestaurantBus(store ListRestaurantStore, like LikeRestaurantStore) *ListRestaurantBus {
	return &ListRestaurantBus{
		store: store,
		like:  like,
	}
}

func (l *ListRestaurantBus) List(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	rs, err := l.store.ListDataWithCondition(context, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.Entity, err)
	}
	ids := make([]int, len(rs))
	for i := range ids {
		ids[i] = rs[i].Id
	}
	likeMap, err := l.like.GetRestaurantLikes(context, ids)
	if err != nil {
		log.Println(err)
		return rs, nil
	}
	for i, item := range rs {
		rs[i].LikedCount = likeMap[item.Id]
	}
	return rs, nil
}
