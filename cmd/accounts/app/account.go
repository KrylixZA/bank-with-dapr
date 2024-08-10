package app

type account struct {
	AccountID   string      `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	Balance     float64     `json:"balance"`
}
