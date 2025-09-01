package model

type Budget struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

type BudgetStore struct {
	Budgets []Budget `json:"budgets"`
}
