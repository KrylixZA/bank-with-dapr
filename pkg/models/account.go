package models

import "github.com/KrylixZA/bank-with-dapr/pkg/enums"

type Account struct {
	AccountID   string            `json:"accountId"`
	AccountType enums.AccountType `json:"accountType"`
	Balance     float64           `json:"balance"`
}
