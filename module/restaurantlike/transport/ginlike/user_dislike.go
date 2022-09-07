package ginlike

import (
	"food/common"
	"food/component/appctx"
	restaurantlikebusiness "food/module/restaurantlike/business"
	restaurantlikestorage "food/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisLike(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := ctx.MustGet("user").(common.Requester)
		store := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
		dis := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
		bus := restaurantlikebusiness.NewUserDisLikeRestaurantBus(store, dis)
		if err := bus.DisLike(ctx.Request.Context(), res.GetUid(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("success"))
	}
}
