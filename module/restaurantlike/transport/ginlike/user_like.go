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

func UserLike(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := ctx.MustGet("user").(common.Requester)
		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       res.GetUid(),
		}
		store := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
		like := restaurantlikestorage.NewSqlStore(appctx.GetMysqlConnection())
		bus := restaurantlikebusiness.NewUserLikeRestaurantBus(store, like)
		if err := bus.Like(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("success"))
	}
}
