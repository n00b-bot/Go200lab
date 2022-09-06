package ginrestaurant

import (
	"food/common"
	"food/component/appctx"
	restaurantsbusiness "food/module/restaurant/business"
	restaurantmodel "food/module/restaurant/model"
	restaurantstorage "food/module/restaurant/storage"
	usermodel "food/module/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		user := c.MustGet("user").(*usermodel.User)
		data.UserId = user.Id
		db := appCtx.GetMysqlConnection()
		store := restaurantstorage.NewSqlStore(db)

		bus := restaurantsbusiness.NewCreateRestaurantBus(store)
		if err := bus.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.GenUID(common.RestaurantType)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
