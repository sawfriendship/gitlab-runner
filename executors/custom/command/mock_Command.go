// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package command

import mock "github.com/stretchr/testify/mock"

// MockCommand is an autogenerated mock type for the Command type
type MockCommand struct {
	mock.Mock
}

// Run provides a mock function with given fields:
func (_m *MockCommand) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
