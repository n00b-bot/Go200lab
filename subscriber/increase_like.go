package subscriber

import (
	"context"
	"food/component/appctx"
	"food/pubsub"
	restaurantlikestorage "food/module/restaurantlike/storage"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

// func IncreaseLike(appctx appctx.AppContext, ctx context.Context) {
// 	c, _ := appctx.GetPubSub().Subscribe(ctx, pubsub.Topic(common.UserLike))
//
// 	go func() {
// 		defer middleware.Recover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			store.UpLike(ctx, likeData.GetRestaurantId())
// 			print("liked")
// 		}
// 	}()
// }

func UpLike(appctx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Like",
		Handler: func(ctx context.Context, msg *pubsub.Message) error {
			likeData := msg.Data().(HasRestaurantId)
			store := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
			return store.UpLike(ctx, likeData.GetRestaurantId())
		},
	}
}
