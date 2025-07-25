// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"
	lib "pkg.formatio/lib"
)

// MockINetworkManager is an autogenerated mock type for the INetworkManager type
type MockINetworkManager struct {
	mock.Mock
}

// CreateNetwork provides a mock function with given fields: _a0
func (_m *MockINetworkManager) CreateNetwork(_a0 lib.CreateNetworkArgs) (*lib.CreateNetworkResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateNetwork")
	}

	var r0 *lib.CreateNetworkResult
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.CreateNetworkArgs) (*lib.CreateNetworkResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(lib.CreateNetworkArgs) *lib.CreateNetworkResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lib.CreateNetworkResult)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.CreateNetworkArgs) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetwork provides a mock function with given fields: _a0
func (_m *MockINetworkManager) DeleteNetwork(_a0 lib.DeleteNetworkArgs) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNetwork")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(lib.DeleteNetworkArgs) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNetwork provides a mock function with given fields: _a0
func (_m *MockINetworkManager) GetNetwork(_a0 lib.GetNetworkArgs) (*lib.CreateNetworkResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetNetwork")
	}

	var r0 *lib.CreateNetworkResult
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.GetNetworkArgs) (*lib.CreateNetworkResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(lib.GetNetworkArgs) *lib.CreateNetworkResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lib.CreateNetworkResult)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.GetNetworkArgs) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields: _a0
func (_m *MockINetworkManager) ListNetworks(_a0 lib.ListNetworkArgs) ([]lib.CreateNetworkResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ListNetworks")
	}

	var r0 []lib.CreateNetworkResult
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.ListNetworkArgs) ([]lib.CreateNetworkResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(lib.ListNetworkArgs) []lib.CreateNetworkResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]lib.CreateNetworkResult)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.ListNetworkArgs) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNetwork provides a mock function with given fields: _a0
func (_m *MockINetworkManager) UpdateNetwork(_a0 lib.UpdateNetworkArgs) (*lib.CreateNetworkResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNetwork")
	}

	var r0 *lib.CreateNetworkResult
	var r1 error
	if rf, ok := ret.Get(0).(func(lib.UpdateNetworkArgs) (*lib.CreateNetworkResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(lib.UpdateNetworkArgs) *lib.CreateNetworkResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lib.CreateNetworkResult)
		}
	}

	if rf, ok := ret.Get(1).(func(lib.UpdateNetworkArgs) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// convertToContainerPort provides a mock function with given fields: _a0
func (_m *MockINetworkManager) convertToContainerPort(_a0 lib.NetworkPort) v1.ContainerPort {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for convertToContainerPort")
	}

	var r0 v1.ContainerPort
	if rf, ok := ret.Get(0).(func(lib.NetworkPort) v1.ContainerPort); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(v1.ContainerPort)
	}

	return r0
}

// convertToServicePort provides a mock function with given fields: _a0
func (_m *MockINetworkManager) convertToServicePort(_a0 lib.NetworkPort) v1.ServicePort {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for convertToServicePort")
	}

	var r0 v1.ServicePort
	if rf, ok := ret.Get(0).(func(lib.NetworkPort) v1.ServicePort); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(v1.ServicePort)
	}

	return r0
}

// NewMockINetworkManager creates a new instance of MockINetworkManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockINetworkManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockINetworkManager {
	mock := &MockINetworkManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
