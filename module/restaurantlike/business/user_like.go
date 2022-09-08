package restaurantlikebusiness

import (
	"context"
	"food/common"
	"food/component/asyncjob"
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
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return u.uplike.UpLike(ctx, data.RestaurantId)
	})
	asyncjob.NewManager(true, j).Run(ctx)
	return nil
}
