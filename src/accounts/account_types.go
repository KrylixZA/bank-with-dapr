package accounts

type AccountType int

const (
	Cheque AccountType = iota
	Savings
	Credit
	Business
)
