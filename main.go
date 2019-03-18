// filename main.go
// author Yoshiya Ito <myon53@gmail.com>
// date 2019-03-13
package main

import (
	"fmt"

	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/funnyface-inc/go-example/db"
	"github.com/funnyface-inc/go-example/middleware"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	viper.BindEnv("ENV")
	switch viper.Get("ENV") {
	case "development":
		viper.SetConfigName("development")
	case "stging":
		viper.SetConfigName("stging")
	case "production":
		viper.SetConfigName("production")
	default:
		viper.SetConfigName("development")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("[VIPER config error]", err)
	}

	// Database Settings
	db.SetupDB()
	db.SetupRedis()
}

func main() {
	api := gin.New()
	api.Use(gin.Logger())
	api.Use(middleware.Transaction())
	api.Use(middleware.Handle500())
	api.Use(middleware.RedisPooling())
	api.Use(gin.Recovery())
	//api.NoRoute(middleware.NotFound)
	//api.NoMethod(middleware.NoMethod)
	//router.V1(api)
	// TODO Gracefull Restart
	api.Run(":8000")
}
