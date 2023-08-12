package rest

import (
	"github.com/J-Obog/paidoff/data"
	"github.com/stretchr/testify/suite"
)

/*
	type AccountUpdateBody struct {
		Name string `json:"name"`
	}
*/
type SerializationTestSuite struct {
	suite.Suite
}

func (s *SerializationTestSuite) TestSerdesJSONBody() {
	testCases := []struct {
		jsons string
		obj   any
	}{
		{
			jsons: `{"categoryId": "cat-123", "projected": 123.45}`,
			obj:   BudgetUpdateBody{CategoryId: "cat-123", Projected: 123.45},
		},
		{
			jsons: `{"categoryId": "cat-123", "projected": 123.45, "month": 8, "year": 2011}`,
			obj:   BudgetCreateBody{CategoryId: "cat-123", Projected: 123.45, Month: 8, Year: 2011},
		},
		{
			jsons: `{"categoryId": null, "note": null, "type": "EXPENSE", "amount": 1234.56, "month": 7, "day": 2, "year": 2022}`,
			obj: TransactionCreateBody{
				CategoryId: nil,
				Note:       nil,
				Type:       data.BudgetType_Expense,
				Amount:     1234.56,
				Month:      7,
				Day:        2,
				Year:       2022,
			},
		},
		{
			jsons: `{"categoryId": null, "note": null, "type": "EXPENSE", "amount": 1234.56, "month": 7, "day": 2, "year": 2022}`,
			obj: TransactionUpdateBody{
				CategoryId: nil,
				Note:       nil,
				Type:       data.BudgetType_Expense,
				Amount:     1234.56,
				Month:      7,
				Day:        2,
				Year:       2022,
			},
		},
		{
			jsons: `{"name": "foobar", "color": 1005}`,
			obj: CategoryCreateBody{
				Name:  "foobar",
				Color: 1005,
			},
		},
		{
			jsons: `{"name": "foobar", "color": 1005}`,
			obj: CategoryUpdateBody{
				Name:  "foobar",
				Color: 1005,
			},
		},
		{
			jsons: `{"name": "John Doe"}`,
			obj: AccountUpdateBody{
				Name: "John Doe",
			},
		},
	}

}
