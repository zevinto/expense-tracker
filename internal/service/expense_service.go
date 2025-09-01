package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/zevinto/expense-tracker/internal/model"
	"github.com/zevinto/expense-tracker/internal/repository"
)

func AddExpense(description string, amountStr string) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount < 0 {
		fmt.Println("Invalid amount: must be a non-negative number")
		return
	}
	store, err := repository.ListExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	nextID := 1
	for _, e := range store.Expenses {
		if e.ID >= nextID {
			nextID = e.ID + 1
		}
	}

	expense := model.Expense{
		ID:          nextID,
		Date:        time.Now().Format("2006-01-02"),
		Description: description,
		Amount:      amount,
	}
	store.Expenses = append(store.Expenses, expense)
	if err := repository.SaveExpenses(store); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Expense added successfully(ID: %d)\n", nextID)
}

func UpdateExpense(idStr string, description string, amountStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	var amount *float64
	if amountStr != "" {
		amt, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amt < 0 {
			fmt.Println("Invalid amount: must be a non-negative number")
			return
		}
		amount = &amt
	}
	if description == "" && amount == nil {
		fmt.Println("Nothing to update. Provide a new description or amount.")
		return
	}
	store, err := repository.ListExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, expense := range store.Expenses {
		if expense.ID == id {
			if description != "" {
				store.Expenses[i].Description = description
			}
			if amount != nil {
				store.Expenses[i].Amount = *amount
			}
			if err := repository.SaveExpenses(store); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Expense updated successfully")
			return
		}
	}
	fmt.Println("Expense not found")
}

func DeleteExpense(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	store, err := repository.ListExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, expense := range store.Expenses {
		if expense.ID == id {
			store.Expenses = append(store.Expenses[:i], store.Expenses[i+1:]...)
			err := repository.SaveExpenses(store)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Printf("expense with ID %d not found", id)
}

func ListExpenses() {
	store, err := repository.ListExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(store.Expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}
	fmt.Println("ID\tDate\tDescription\tAmount")
	for _, expense := range store.Expenses {
		fmt.Printf("%d\t%s\t%s\t$%.2f\n", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
}

func SumExpensesByMonth(store model.ExpenseStore, year, month int) float64 {
	total := 0.0
	for _, expense := range store.Expenses {
		date, err := time.Parse("2006-01-02", expense.Date)
		if err == nil && date.Year() == year && int(date.Month()) == month {
			total += expense.Amount
		}
	}
	return total
}

func Summary(monthStr string) {
	store, err := repository.ListExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	currentYear := time.Now().Year()
	if monthStr != "" {
		month, err := strconv.Atoi(monthStr)
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid month: must be between 1 and 12")
			return
		}
		total := 0.0
		for _, expense := range store.Expenses {
			date, err := time.Parse("2006-01-02", expense.Date)
			if err == nil && date.Year() == currentYear && int(date.Month()) == month {
				total += expense.Amount
			}
		}
		fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(month).String(), total)
	} else {
		total := 0.0
		for _, expense := range store.Expenses {
			total += expense.Amount
		}
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

func ExportCSV() {
	expenseStore, err := repository.ListExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		return
	}
	file, err := os.Create("expenses_export.csv")
	if err != nil {
		fmt.Printf("Error creating CSV file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"ID", "Date", "Description", "Amount"}); err != nil {
		fmt.Printf("Error writing CSV header: %v\n", err)
		return
	}

	// Write Data
	for _, expense := range expenseStore.Expenses {
		record := []string{
			strconv.Itoa(expense.ID),
			expense.Date,
			expense.Description,
			strconv.FormatFloat(expense.Amount, 'f', -1, 64),
		}
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing CSV record: %v\n", err)
			return
		}
	}
	fmt.Println("Expenses exported successfully to expenses_export.csv")
}
