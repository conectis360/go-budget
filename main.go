package main

import (
	"fmt"
	"sync"
	"time"
)

type Expense struct {
	Category string
	Amount   float64
	Date     time.Time
}

type User struct {
	Name     string
	Expenses []Expense
}

func recordExpense(user *User, category string, amount float64, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when finished
	expense := Expense{
		Category: category,
		Amount:   amount,
		Date:     time.Now(),
	}
	user.Expenses = append(user.Expenses, expense)
	fmt.Printf("Recorded expense: %s - $%.2f\n", category, amount)
}

func main() {
	var wg sync.WaitGroup
	user := User{Name: "Alice"}

	// Simulate recording expenses concurrently using Go routines
	wg.Add(3)
	go recordExpense(&user, "Groceries", 50.00, &wg)
	go recordExpense(&user, "Rent", 1200.00, &wg)
	go recordExpense(&user, "Utilities", 100.00, &wg)

	wg.Wait() // Wait for all goroutines to finish

	// Print the recorded expenses
	fmt.Println("Expenses recorded for user:", user.Name)
	for _, expense := range user.Expenses {
		fmt.Printf("%s: $%.2f on %s\n", expense.Category, expense.Amount, expense.Date.Format("2006-01-02"))
	}
}
