package data

import "bank_backend/models"

var Accounts = map[string]*models.Account{
	"Alice": {
		Owner:   "Alice",
		Balance: 1000.0,
		Transactions: []models.Transaction{
			{Amount: 1000.0, Description: "Initial deposit"},
		},
	},
	"Bob": {
		Owner:   "Bob",
		Balance: 500.0,
		Transactions: []models.Transaction{
			{Amount: 500.0, Description: "Initial deposit"},
		},
	},
}
