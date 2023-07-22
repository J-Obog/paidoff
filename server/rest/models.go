package rest

import (
	"github.com/J-Obog/paidoff/data"
)

type BudgetQuery struct {
	Month *int `json:"month"`
	Year  *int `json:"year"`
}

type TransactionQuery struct {
	CreatedBefore *int64   `json:"createdBefore"`
	CreatedAfter  *int64   `json:"createdAfter"`
	AmountGte     *float64 `json:"amountGte"`
	AmountLte     *float64 `json:"amountLte"`
}

type AccountSetBody struct {
	Name string `json:"name"`
}

type AccountUpdateBody struct {
	AccountSetBody
}

type AccountCreateBody struct {
	AccountSetBody
}

type CategorySetBody struct {
	Name  string `json:"name"`
	Color uint   `json:"color"`
}

type CategoryUpdateBody struct {
	CategorySetBody
}

type CategoryCreateBody struct {
	CategorySetBody
}

type TransactionSetBody struct {
	CategoryId *string         `json:"categoryId"`
	Note       *string         `json:"note"`
	Type       data.BudgetType `json:"budgetType"`
	Amount     float64         `json:"amount"`
	Month      int             `json:"month"`
	Day        int             `json:"day"`
	Year       int             `json:"year"`
}

type TransactionUpdateBody struct {
	TransactionSetBody
}

type TransactionCreateBody struct {
	TransactionSetBody
}

type BudgetSetBody struct {
	CategoryId string  `json:"categoryId"`
	Projected  float64 `json:"projected"`
}

type BudgetCreateBody struct {
	BudgetSetBody
	Month int `json:"month"`
	Year  int `json:"year"`
}

type BudgetUpdateBody struct {
	BudgetSetBody
}
