package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zevinto/expense-tracker/internal/model"
)

const BudgetFile = "budgets.json"

func LoadBudgets() (model.BudgetStore, error) {
	var store model.BudgetStore
	if _, err := os.Stat(BudgetFile); os.IsNotExist(err) {
		return model.BudgetStore{Budgets: []model.Budget{}}, nil
	}

	data, err := os.ReadFile(BudgetFile)
	if err != nil {
		return store, fmt.Errorf("failed to read task file: %w", err)
	}
	err = json.Unmarshal(data, &store)
	if err != nil {
		return store, fmt.Errorf("failed to parse task file: %w", err)
	}
	return store, nil
}

func SaveBudgets(store model.BudgetStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	err = os.WriteFile(BudgetFile, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}
	return nil
}
