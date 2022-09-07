package restaurantlikebusiness

import (
	"context"
	"food/common"
	restaurantlikemodel "food/module/restaurantlike/model"
)

type listUsersLikeStore interface {
	GetUsersLikeRestaurant(ctx context.Context, conditions map[string]interface{}, filter *restaurantlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error)
}

type ListUsersLikeBus struct {
	store listUsersLikeStore
}

func NewListUsersLikeBus(store listUsersLikeStore) *ListUsersLikeBus {
	return &ListUsersLikeBus{
		store: store,
	}
}

func (l *ListUsersLikeBus) ListUsers(ctx context.Context, filter *restaurantlikemodel.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := l.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}
	return users, nil
}
