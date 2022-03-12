package main

import (
	"users/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	router := gin.New()
	handler := users.NewHandler()
	handler.Register(router)
	port := viper.GetString("server.port")
	router.Run(":" + port)
}
