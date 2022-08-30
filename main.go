package main

import (
	"food/component/appctx"
	"food/component/uploadprovider"
	"food/middleware"
	"food/module/restaurant/transport/ginrestaurant"
	"food/module/upload/transport/ginupload"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	s3BucketName := viper.Get("s3BucketName").(string)
	s3Region := viper.Get("s3Region").(string)
	s3APIKey := viper.Get("s3APIKey").(string)
	s3SecretKey := viper.Get("s3SecretKey").(string)
	s3Domain := viper.Get("s3Domain").(string)
	//secrectKey := viper.Get("SYSTEM_SECRET")

	s3 := uploadprovider.NewS3(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appCtx := appctx.NewAppCtx(db, s3)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	v1.POST("/upload", ginupload.UploadImage(appCtx))
	res := v1.Group("/restaurant")
	res.POST("", ginrestaurant.CreateRestaurant(appCtx))
	res.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	res.GET("", ginrestaurant.ListRestaurant(appCtx))
	r.Run()
}
