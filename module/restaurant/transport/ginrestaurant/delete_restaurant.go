package ginrestaurant

import (
	"food/common"
	"food/component/appctx"
	restaurantsbusiness "food/module/restaurant/business"
	restaurantstorage "food/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMysqlConnection()
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStore(db)
		bus := restaurantsbusiness.NewDeleteRestaurantBus(store)
		if err := bus.Delete(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
