package model

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type ExpenseStore struct {
	Expenses []Expense `json:"expenses"`
}
