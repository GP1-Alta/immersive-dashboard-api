// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	status "immersive-dashboard/features/status"

	mock "github.com/stretchr/testify/mock"
)

// StatusData is an autogenerated mock type for the StatusDataInterface type
type StatusData struct {
	mock.Mock
}

// SelectAll provides a mock function with given fields:
func (_m *StatusData) SelectAll() ([]status.Core, error) {
	ret := _m.Called()

	var r0 []status.Core
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]status.Core, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []status.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]status.Core)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStatusData interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusData creates a new instance of StatusData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusData(t mockConstructorTestingTNewStatusData) *StatusData {
	mock := &StatusData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
