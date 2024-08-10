package main

import (
	accounts "github.com/KrylixZA/bank-with-dapr/cmd/accounts/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/accounts/:userId", accounts.GetAccountsForUser)
	router.POST("/accounts/:userId", accounts.CreateAccountForUser)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
