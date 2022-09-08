package restaurantlikebusiness

import (
	"context"
	"food/common"
	"food/component/asyncjob"
	restaurantlikemodel "food/module/restaurantlike/model"
)

type UserDisLikeRestaurantStore interface {
	Delete(context.Context, int, int) error
}

type DownLikeStore interface {
	DownLike(ctx context.Context, id int) error
}

type UserDisLikeRestaurantBus struct {
	store    UserDisLikeRestaurantStore
	downlike DownLikeStore
}

func NewUserDisLikeRestaurantBus(store UserDisLikeRestaurantStore, like DownLikeStore) *UserDisLikeRestaurantBus {
	return &UserDisLikeRestaurantBus{
		store:    store,
		downlike: like,
	}
}

func (u *UserDisLikeRestaurantBus) DisLike(ctx context.Context, userid int, resid int) error {
	if err := u.store.Delete(ctx, userid, resid); err != nil {
		return common.ErrCannotDeleteEntity(restaurantlikemodel.EntityName, err)
	}
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return u.downlike.DownLike(ctx, resid)
	})
	asyncjob.NewManager(true, j).Run(ctx)

	return nil
}
