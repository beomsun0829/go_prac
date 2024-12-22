package handlers

import (
	"bank_backend/data"
	"bank_backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetAccounts(c *fiber.Ctx) error {
	accounts := []models.Account{}

	for _, account := range data.Accounts {
		accounts = append(accounts, *account)
	}

	return c.JSON(accounts)
}

func GetAccount(c *fiber.Ctx) error {
	owner := c.Params("owner")

	account, exists := data.Accounts[owner]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}
	return c.JSON(account)
}

type TransactionRequest struct {
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func CreateTransaction(c *fiber.Ctx) error {
	var req TransactionRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.From == "" || req.To == "" || req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or invalid parameters",
		})
	}

	fromAccount, fromExists := data.Accounts[req.From]
	if !fromExists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sender account not found",
		})
	}

	toAccount, toExists := data.Accounts[req.To]
	if !toExists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Recipient account not found",
		})
	}

	if fromAccount.Balance < req.Amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Insufficient funds",
		})
	}

	// Apply Transaction
	fromAccount.Balance -= req.Amount
	toAccount.Balance += req.Amount

	fromTransaction := models.Transaction{
		Amount:      -req.Amount,
		Description: "To " + req.To + ": " + req.Description,
	}

	toTransaction := models.Transaction{
		Amount:      req.Amount,
		Description: "From " + req.From + ": " + req.Description,
	}

	fromAccount.Transactions = append(fromAccount.Transactions, fromTransaction)
	toAccount.Transactions = append(toAccount.Transactions, toTransaction)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Transaction successful",
		"from":        req.From,
		"to":          req.To,
		"amount":      req.Amount,
		"description": req.Description,
	})
}
