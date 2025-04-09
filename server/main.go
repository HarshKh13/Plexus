package main

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("users", api.CreateUser)
	router.GET("users", api.GetAllUsers)
	router.Run(":8080")
}
