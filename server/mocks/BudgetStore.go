// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	data "github.com/J-Obog/paidoff/data"
	mock "github.com/stretchr/testify/mock"
)

// BudgetStore is an autogenerated mock type for the BudgetStore type
type BudgetStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id, accountId
func (_m *BudgetStore) Delete(id string, accountId string) (bool, error) {
	ret := _m.Called(id, accountId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(id, accountId)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(id, accountId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields:
func (_m *BudgetStore) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id, accountId
func (_m *BudgetStore) Get(id string, accountId string) (*data.Budget, error) {
	ret := _m.Called(id, accountId)

	var r0 *data.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*data.Budget, error)); ok {
		return rf(id, accountId)
	}
	if rf, ok := ret.Get(0).(func(string, string) *data.Budget); ok {
		r0 = rf(id, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Budget)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBy provides a mock function with given fields: accountId, filter
func (_m *BudgetStore) GetBy(accountId string, filter data.BudgetFilter) ([]data.Budget, error) {
	ret := _m.Called(accountId, filter)

	var r0 []data.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(string, data.BudgetFilter) ([]data.Budget, error)); ok {
		return rf(accountId, filter)
	}
	if rf, ok := ret.Get(0).(func(string, data.BudgetFilter) []data.Budget); ok {
		r0 = rf(accountId, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]data.Budget)
		}
	}

	if rf, ok := ret.Get(1).(func(string, data.BudgetFilter) error); ok {
		r1 = rf(accountId, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCategory provides a mock function with given fields: accountId, categoryId
func (_m *BudgetStore) GetByCategory(accountId string, categoryId string) ([]data.Budget, error) {
	ret := _m.Called(accountId, categoryId)

	var r0 []data.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]data.Budget, error)); ok {
		return rf(accountId, categoryId)
	}
	if rf, ok := ret.Get(0).(func(string, string) []data.Budget); ok {
		r0 = rf(accountId, categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]data.Budget)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountId, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPeriodCategory provides a mock function with given fields: accountId, categoryId, month, year
func (_m *BudgetStore) GetByPeriodCategory(accountId string, categoryId string, month int, year int) (*data.Budget, error) {
	ret := _m.Called(accountId, categoryId, month, year)

	var r0 *data.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) (*data.Budget, error)); ok {
		return rf(accountId, categoryId, month, year)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) *data.Budget); ok {
		r0 = rf(accountId, categoryId, month, year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Budget)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(accountId, categoryId, month, year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: budget
func (_m *BudgetStore) Insert(budget data.Budget) error {
	ret := _m.Called(budget)

	var r0 error
	if rf, ok := ret.Get(0).(func(data.Budget) error); ok {
		r0 = rf(budget)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, accountId, update, timestamp
func (_m *BudgetStore) Update(id string, accountId string, update data.BudgetUpdate, timestamp int64) (bool, error) {
	ret := _m.Called(id, accountId, update, timestamp)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, data.BudgetUpdate, int64) (bool, error)); ok {
		return rf(id, accountId, update, timestamp)
	}
	if rf, ok := ret.Get(0).(func(string, string, data.BudgetUpdate, int64) bool); ok {
		r0 = rf(id, accountId, update, timestamp)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, data.BudgetUpdate, int64) error); ok {
		r1 = rf(id, accountId, update, timestamp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBudgetStore creates a new instance of BudgetStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBudgetStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *BudgetStore {
	mock := &BudgetStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
