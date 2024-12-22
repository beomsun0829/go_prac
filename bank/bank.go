package main

import (
	"fmt"
	"strings"
)

type Transaction struct {
	Amount      float64
	Description string
}

type Account struct {
	Owner        string
	Balance      float64
	Transactions []Transaction
}

func (a *Account) Deposit(amount float64, desc string) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	a.Balance += amount
	a.Transactions = append(a.Transactions, Transaction{amount, desc})
	fmt.Printf("Deposited %.2f to %s's account.\n", amount, a.Owner)
	fmt.Printf("New balance: %.2f\n", a.Balance)
}

func (a *Account) Withdraw(amount float64, desc string) {
	if amount <= 0 {
		fmt.Println("Withdraw amount must be positive")
		return
	}
	if amount > a.Balance {
		fmt.Println("Insufficient funds")
		return
	}
	a.Balance -= amount
	a.Transactions = append(a.Transactions, Transaction{-amount, desc})
	fmt.Printf("Withdraw %.2f from %s's account.\n", amount, a.Owner)
	fmt.Printf("New balance: %.2f\n", a.Balance)
}

func (a *Account) PrintStatement() {
	fmt.Printf("Account Statement for %s\n", a.Owner)
	fmt.Println(strings.Repeat("-", 30))

	for _, t := range a.Transactions {
		fmt.Printf("%-20s %+10.2f\n", t.Description, t.Amount)
	}

	fmt.Printf("Current Balance: %.2f\n", a.Balance)
	fmt.Println(strings.Repeat("-", 30))
}

type Bank struct {
	Accounts map[string]*Account
}

func (b *Bank) CreateAccount(owner string) {
	_, exists := b.Accounts[owner]
	if exists {
		fmt.Println("Account already exists")
		return
	}
	b.Accounts[owner] = &Account{owner, 0.0, []Transaction{}}
	fmt.Printf("Account created: %s\n", owner)
}

func (b *Bank) GetAccount(owner string) *Account {
	account, exists := b.Accounts[owner]
	if !exists {
		fmt.Println("Account not found")
		return nil
	}
	return account
}

func main() {
	bank := Bank{Accounts: make(map[string]*Account)}

	bank.CreateAccount("Alice")
	bank.CreateAccount("Bob")

	alice := bank.GetAccount("Alice")
	if alice != nil {
		alice.Deposit(1000, "Initial Deposit")
		alice.Withdraw(200, "Groceries")
		alice.Deposit(500, "Freelance Payment")
		alice.PrintStatement()
	}

	bob := bank.GetAccount("Bob")
	if bob != nil {
		bob.Deposit(1500, "Salary")
		bob.Withdraw(400, "Rent")
		bob.Deposit(300, "Bonus")
		bob.Withdraw(2000, "Vacation")
		bob.PrintStatement()
	}

	bank.GetAccount("Charlie")
}
