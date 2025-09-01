package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zevinto/expense-tracker/internal/model"
)

const ExpenseFile = "expenses.json"

func ListExpenses() (model.ExpenseStore, error) {
	var store model.ExpenseStore
	if _, err := os.Stat(ExpenseFile); os.IsNotExist(err) {
		return model.ExpenseStore{Expenses: []model.Expense{}}, nil
	}

	data, err := os.ReadFile(ExpenseFile)
	if err != nil {
		return store, fmt.Errorf("failed to read task file: %w", err)
	}
	err = json.Unmarshal(data, &store)
	if err != nil {
		return store, fmt.Errorf("failed to parse task file: %w", err)
	}
	return store, nil
}

func SaveExpenses(store model.ExpenseStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	err = os.WriteFile(ExpenseFile, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}
	return nil
}
