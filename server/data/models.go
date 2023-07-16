package data

type BudgetType string

const (
	BudgetType_Income  BudgetType = "Income"
	BudgetType_Expense BudgetType = "Expense"
)

type Date struct {
	Month int
	Day   int
	Year  int
}

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	IsActivated bool  `json:"isActivated"`
	IsDeleted   bool  `json:"isDeleted"`
	CreatedAt   int64 `json:"createdAt"`
	UpdatedAt   int64 `json:"updatedAt"`
}

type BudgetList []Budget

func (arr BudgetList) First() *Budget {
	if len(arr) > 0 {
		return &arr[0]
	}
	return nil
}

type Budget struct {
	Id         string  `json:"id"`
	AccountId  string  `json:"accountId"`
	CategoryId string  `json:"categoryId"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
	Projected  float64 `json:"projected"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

type BudgetFilter struct {
	AccountId  *string
	CategoryId *string
	Month      *int
	Year       *int
}

type BudgetMaterialized struct {
	Actual float64 `json:"actual"`
}

type CategoryList []Category

func (arr CategoryList) First() *Category {
	if len(arr) > 0 {
		return &arr[0]
	}
	return nil
}

type Category struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	Color     uint   `json:"color"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

type CategoryFilter struct {
	AccountId *string
	Name      *string
}

type TransactionList []Transaction

func (arr TransactionList) First() *Transaction {
	if len(arr) > 0 {
		return &arr[0]
	}
	return nil
}

type Transaction struct {
	Id         string     `json:"id"`
	AccountId  string     `json:"accountId"`
	CategoryId *string    `json:"categoryId"`
	Note       *string    `json:"note"`
	Type       BudgetType `json:"budgetType"`
	Amount     float64    `json:"amount"`
	Month      int        `json:"month"`
	Day        int        `json:"day"`
	Year       int        `json:"year"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
}

type TransactionFilter struct {
	Before      *Date
	After       *Date
	GreaterThan *float64
	LessThan    *float64
	CategoryId  *string
	AccountId   *string
}
