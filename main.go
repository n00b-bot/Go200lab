package main

import (
	"food/module/restaurant/transport/ginrestaurant"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:nothing@tcp(127.0.0.1:3306)/food?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	r := gin.Default()

	v1 := r.Group("/v1")
	res := v1.Group("/restaurant")
	res.POST("", ginrestaurant.CreateRestaurant(db))
	res.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	r.Run()
}
