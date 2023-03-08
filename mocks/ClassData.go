// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	classes "immersive-dashboard/features/classes"

	mock "github.com/stretchr/testify/mock"
)

// ClassData is an autogenerated mock type for the ClassDataInterface type
type ClassData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: data, id
func (_m *ClassData) Delete(data classes.Core, id uint) error {
	ret := _m.Called(data, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(classes.Core, uint) error); ok {
		r0 = rf(data, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *ClassData) Insert(input classes.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(classes.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAll provides a mock function with given fields:
func (_m *ClassData) ListAll() ([]classes.Core, error) {
	ret := _m.Called()

	var r0 []classes.Core
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]classes.Core, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []classes.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]classes.Core)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAll provides a mock function with given fields: limit, offset, name
func (_m *ClassData) SelectAll(limit int, offset int, name string) ([]classes.Core, error) {
	ret := _m.Called(limit, offset, name)

	var r0 []classes.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]classes.Core, error)); ok {
		return rf(limit, offset, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []classes.Core); ok {
		r0 = rf(limit, offset, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]classes.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *ClassData) Update(input classes.Core, id uint) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(classes.Core, uint) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClassData interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassData creates a new instance of ClassData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassData(t mockConstructorTestingTNewClassData) *ClassData {
	mock := &ClassData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}