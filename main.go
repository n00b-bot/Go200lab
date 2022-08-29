package main

import (
	"food/component/appctx"
	"food/middleware"
	"food/module/restaurant/transport/ginrestaurant"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "root:nothing@tcp(127.0.0.1:3306)/food?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		return
	}
	appCtx := appctx.NewAppCtx(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	res := v1.Group("/restaurant")
	res.POST("", ginrestaurant.CreateRestaurant(appCtx))
	res.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	res.GET("", ginrestaurant.ListRestaurant(appCtx))
	r.Run()
}
