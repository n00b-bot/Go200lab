package ginrestaurant

import (
	restaurantsbusiness "food/module/restaurant/business"
	restaurantstorage "food/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id int
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStore(db)
		bus := restaurantsbusiness.NewDeleteRestaurantBus(store)
		if err := bus.Delete(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	}
}
