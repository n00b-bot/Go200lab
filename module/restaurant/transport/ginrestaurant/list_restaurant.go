package ginrestaurant

import (
	"food/common"
	"food/component/appctx"
	restaurantsbusiness "food/module/restaurant/business"
	restaurantmodel "food/module/restaurant/model"
	restaurantstorage "food/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rs []restaurantmodel.Restaurant
		var filter restaurantmodel.Filter
		var paging common.Paging
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Status = []int{1}
		paging.Fullfill()
		db := appCtx.GetMysqlConnection()
		store := restaurantstorage.NewSqlStore(db)
		bus := restaurantsbusiness.NewListRestaurantBus(store)
		rs, err := bus.List(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range rs {
			rs[i].GenUID(common.RestaurantType)

		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(rs, paging, filter))
	}

}
