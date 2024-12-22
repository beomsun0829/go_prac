package models

type Transaction struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type Account struct {
	Owner        string        `json:"owner"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}
