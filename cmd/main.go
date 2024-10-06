package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userGroup := router.Group("/v1/users")
	userServer, err := InitializeUserServer()
	if err != nil {
		panic(err)
	}
	userServer.RegisterHandlers(userGroup)
	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}


