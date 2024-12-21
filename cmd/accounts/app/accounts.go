package app

import (
	"encoding/json"
	"github.com/KrylixZA/bank-with-dapr/pkg/models"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

const state = "accounts-state"
const validationFailureErrorMessage = "model validation failure while attempting to map request to type"
const notFoundErrorMessage = "resource not found"
const internalServerErrorMessage = "an environmental, non-specific error has occurred"

func GetAccountsForUser(c *gin.Context) {
	userId := c.Param("userId")

	accounts, err := getAccountsForUser(userId, c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMessage})
		return
	}

	if len(accounts) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": notFoundErrorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, accounts)
}

func CreateAccountForUser(c *gin.Context) {
	userId := c.Param("userId")
	var account models.Account

	accounts, err := getAccountsForUser(userId, c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMessage})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationFailureErrorMessage})
		return
	}

	accounts = append(accounts, account)
	if err := saveAccountsForUser(userId, accounts, c); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMessage})
		return
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

func getAccountsForUser(userId string, c *gin.Context) (accounts []models.Account, err error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	daprStateItem, err := client.GetState(c, state, userId, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(daprStateItem.Value, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func saveAccountsForUser(userId string, accounts []models.Account, c *gin.Context) (err error) {
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}
	defer client.Close()

	accountBytes, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	err = client.SaveState(c, state, userId, accountBytes, nil)
	if err != nil {
		return err
	}

	return nil
}
