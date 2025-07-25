// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	db "pkg.formatio/prisma/db"

	types "pkg.formatio/types"
)

// MockITransactionDao is an autogenerated mock type for the ITransactionDao type
type MockITransactionDao struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: args
func (_m *MockITransactionDao) CreateTransaction(args types.CreateTransactionArgs) (*db.TransactionModel, error) {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 *db.TransactionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(types.CreateTransactionArgs) (*db.TransactionModel, error)); ok {
		return rf(args)
	}
	if rf, ok := ret.Get(0).(func(types.CreateTransactionArgs) *db.TransactionModel); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.TransactionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(types.CreateTransactionArgs) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransaction provides a mock function with given fields: args
func (_m *MockITransactionDao) DeleteTransaction(args types.DeleteTransactionArgs) error {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.DeleteTransactionArgs) error); ok {
		r0 = rf(args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTransaction provides a mock function with given fields: args
func (_m *MockITransactionDao) GetTransaction(args types.GetTransactionArgs) (*db.TransactionModel, error) {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 *db.TransactionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(types.GetTransactionArgs) (*db.TransactionModel, error)); ok {
		return rf(args)
	}
	if rf, ok := ret.Get(0).(func(types.GetTransactionArgs) *db.TransactionModel); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.TransactionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(types.GetTransactionArgs) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransaction provides a mock function with given fields: args
func (_m *MockITransactionDao) ListTransaction(args types.ListTransactionArgs) ([]db.TransactionModel, error) {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for ListTransaction")
	}

	var r0 []db.TransactionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(types.ListTransactionArgs) ([]db.TransactionModel, error)); ok {
		return rf(args)
	}
	if rf, ok := ret.Get(0).(func(types.ListTransactionArgs) []db.TransactionModel); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.TransactionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(types.ListTransactionArgs) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: args
func (_m *MockITransactionDao) UpdateTransaction(args types.UpdateTransactionArgs) (*db.TransactionModel, error) {
	ret := _m.Called(args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransaction")
	}

	var r0 *db.TransactionModel
	var r1 error
	if rf, ok := ret.Get(0).(func(types.UpdateTransactionArgs) (*db.TransactionModel, error)); ok {
		return rf(args)
	}
	if rf, ok := ret.Get(0).(func(types.UpdateTransactionArgs) *db.TransactionModel); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.TransactionModel)
		}
	}

	if rf, ok := ret.Get(1).(func(types.UpdateTransactionArgs) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockITransactionDao creates a new instance of MockITransactionDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockITransactionDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockITransactionDao {
	mock := &MockITransactionDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
