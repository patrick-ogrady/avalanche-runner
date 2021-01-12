// Code generated by mockery. DO NOT EDIT.

package health

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// IsBootstrapped provides a mock function with given fields: chain
func (_m *Client) IsBootstrapped(chain string) (bool, error) {
	ret := _m.Called(chain)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(chain)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsHealthy provides a mock function with given fields:
func (_m *Client) IsHealthy() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
