// Code generated by mockery v2.12.2. DO NOT EDIT.

package cachetype

import (
	"context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockRPC is an autogenerated mock type for the RPC type
type MockRPC struct {
	mock.Mock
}

// RPC provides a mock function with given fields: ctx, method, args, reply
func (_m *MockRPC) RPC(ctx context.Context, method string, args interface{}, reply interface{}) error {
	ret := _m.Called(method, args, reply)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, interface{}) error); ok {
		r0 = rf(method, args, reply)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRPC creates a new instance of MockRPC. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRPC(t testing.TB) *MockRPC {
	mock := &MockRPC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
