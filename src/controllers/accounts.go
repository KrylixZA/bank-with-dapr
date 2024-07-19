package controllers

import (
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

const state = "accounts-state"

func GetAccounts(c *gin.Context) {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	accountId := c.Param("id")

	accounts, err := client.GetState(c, state, accountId, nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "account not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, accounts)
}
