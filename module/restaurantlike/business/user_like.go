package restaurantlikebusiness

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(context.Context, *restaurantlikemodel.Like) error
}

type UserLikeRestaurantBus struct {
	store  UserLikeRestaurantStore
	uplike UpLikeStore
}

type UpLikeStore interface {
	UpLike(ctx context.Context, id int) error
}

func NewUserLikeRestaurantBus(store UserLikeRestaurantStore, like UpLikeStore) *UserLikeRestaurantBus {
	return &UserLikeRestaurantBus{
		store:  store,
		uplike: like,
	}
}

func (u *UserLikeRestaurantBus) Like(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := u.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantlikemodel.EntityName, err)
	}
	u.uplike.UpLike(ctx, data.RestaurantId)
	return nil
}
