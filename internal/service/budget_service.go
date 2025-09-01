package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/zevinto/expense-tracker/internal/model"
	"github.com/zevinto/expense-tracker/internal/repository"
)

func SetBudget(monthStr string, amountStr string) {
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		fmt.Println("Invalid month: must be between 1 and 12")
		return
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount < 0 {
		fmt.Println("Invalid amount: must be a non-negative number")
		return
	}
	budgets, err := repository.LoadBudgets()
	if err != nil {
		fmt.Printf("Error loading budgets: %v\n", err)
		return
	}
	currentYear := time.Now().Year()
	newBudgetStore := model.BudgetStore{
		Budgets: make([]model.Budget, 0, len(budgets.Budgets)),
	}
	found := false
	for _, budget := range budgets.Budgets {
		if budget.Year == currentYear && budget.Month == month {
			newBudgetStore.Budgets = append(newBudgetStore.Budgets, model.Budget{Year: currentYear, Month: month, Amount: amount})
			found = true
		} else {
			newBudgetStore.Budgets = append(newBudgetStore.Budgets, budget)
		}
	}
	if !found {
		newBudgetStore.Budgets = append(newBudgetStore.Budgets, model.Budget{Year: currentYear, Month: month, Amount: amount})
	}
	if err := repository.SaveBudgets(newBudgetStore); err != nil {
		fmt.Printf("Error saving budgets: %v\n", err)
		return
	}

	fmt.Printf("Budget set for %s: $%.2f\n", time.Month(month).String(), amount)

	// Check current expenses against new budget
	expenses, err := repository.ListExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		return
	}
	total := SumExpensesByMonth(expenses, currentYear, month)
	if total > amount {
		fmt.Printf("Warning: Budget of $%.2f for %s exceeded. Current total: $%.2f\n",
			amount, time.Month(month).String(), total)
	}
}
