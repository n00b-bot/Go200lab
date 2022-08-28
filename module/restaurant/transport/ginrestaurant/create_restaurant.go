package ginrestaurant

import (
	restaurantsbusiness "food/module/restaurant/business"
	restaurantmodel "food/module/restaurant/model"
	restaurantstorage "food/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStore(db)
		bus := restaurantsbusiness.NewCreateRestaurantBus(store)
		if err := bus.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
