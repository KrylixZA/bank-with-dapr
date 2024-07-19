package main

import (
	"github.com/KrylixZA/bank-with-dapr/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/accounts/:id", controllers.GetAccounts)

	router.Run("localhost:8080")
}
