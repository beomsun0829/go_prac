package main

import (
	"bank_backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/accounts", handlers.GetAccounts)
	app.Get("/accounts/:owner", handlers.GetAccount)
	app.Post("/transactions", handlers.CreateTransaction)

	app.Listen(":8080")
}
