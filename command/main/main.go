package main

import (
	"context"
	"fmt"
	"users/internal/config"
	"users/internal/users"
	db "users/internal/users/mongodb"
	"users/package/client/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	config, err := config.ViperConfig()
	if err != nil {
		panic(err)
	}

	mongoDBClient, err := mongodb.NewClient(context.Background(),
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.Database,
		config.AuthDB,
	)
	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(mongoDBClient, config.Collection)
	fmt.Println(storage)

	router := gin.New()
	handler := users.NewHandler()
	handler.Register(router)
	port := viper.GetString("server.port")
	router.Run(":" + port)

}
