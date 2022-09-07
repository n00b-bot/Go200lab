package ginlike

import (
	"food/common"
	"food/component/appctx"
	restaurantlikebusiness "food/module/restaurantlike/business"
	restaurantlikemodel "food/module/restaurantlike/model"
	restaurantlikestorage "food/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fullfill()
		store := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
		bus := restaurantlikebusiness.NewListUsersLikeBus(store)
		users, err := bus.ListUsers(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i, _ := range users {
			
			users[i].GenUID(2)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))

	}
}
