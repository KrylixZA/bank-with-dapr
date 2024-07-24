package main

import (
	"github.com/KrylixZA/bank-with-dapr/accounts"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/accounts/:userId", accounts.GetAccountsForUser)
	router.POST("/accounts/:userId", accounts.CreateAccountForUser)

	router.Run("localhost:8080")
}
