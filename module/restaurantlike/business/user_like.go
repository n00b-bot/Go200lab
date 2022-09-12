package restaurantlikebusiness

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
	"food/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(context.Context, *restaurantlikemodel.Like) error
}

type UserLikeRestaurantBus struct {
	store  UserLikeRestaurantStore
	//uplike UpLikeStore
	ps     pubsub.PubSub
}

// type UpLikeStore interface {
// 	UpLike(ctx context.Context, id int) error
// }

func NewUserLikeRestaurantBus(store UserLikeRestaurantStore, ps pubsub.PubSub) *UserLikeRestaurantBus {
	return &UserLikeRestaurantBus{
		store:  store,
		//uplike: like,
		ps:     ps,
	}
}

func (u *UserLikeRestaurantBus) Like(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := u.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantlikemodel.EntityName, err)
	}
	u.ps.Publish(ctx, pubsub.Topic(common.UserLike), pubsub.NewMessage(data))
	return nil
}
