// Code generated by mockery v2.52.2. DO NOT EDIT.

package networking

import (
	mock "github.com/stretchr/testify/mock"
	networking "github.com/wavesplatform/gowaves/pkg/networking"
)

// MockProtocol is an autogenerated mock type for the Protocol type
type MockProtocol struct {
	mock.Mock
}

type MockProtocol_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProtocol) EXPECT() *MockProtocol_Expecter {
	return &MockProtocol_Expecter{mock: &_m.Mock}
}

// EmptyHandshake provides a mock function with no fields
func (_m *MockProtocol) EmptyHandshake() networking.Handshake {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EmptyHandshake")
	}

	var r0 networking.Handshake
	if rf, ok := ret.Get(0).(func() networking.Handshake); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(networking.Handshake)
		}
	}

	return r0
}

// MockProtocol_EmptyHandshake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EmptyHandshake'
type MockProtocol_EmptyHandshake_Call struct {
	*mock.Call
}

// EmptyHandshake is a helper method to define mock.On call
func (_e *MockProtocol_Expecter) EmptyHandshake() *MockProtocol_EmptyHandshake_Call {
	return &MockProtocol_EmptyHandshake_Call{Call: _e.mock.On("EmptyHandshake")}
}

func (_c *MockProtocol_EmptyHandshake_Call) Run(run func()) *MockProtocol_EmptyHandshake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProtocol_EmptyHandshake_Call) Return(_a0 networking.Handshake) *MockProtocol_EmptyHandshake_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProtocol_EmptyHandshake_Call) RunAndReturn(run func() networking.Handshake) *MockProtocol_EmptyHandshake_Call {
	_c.Call.Return(run)
	return _c
}

// EmptyHeader provides a mock function with no fields
func (_m *MockProtocol) EmptyHeader() networking.Header {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EmptyHeader")
	}

	var r0 networking.Header
	if rf, ok := ret.Get(0).(func() networking.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(networking.Header)
		}
	}

	return r0
}

// MockProtocol_EmptyHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EmptyHeader'
type MockProtocol_EmptyHeader_Call struct {
	*mock.Call
}

// EmptyHeader is a helper method to define mock.On call
func (_e *MockProtocol_Expecter) EmptyHeader() *MockProtocol_EmptyHeader_Call {
	return &MockProtocol_EmptyHeader_Call{Call: _e.mock.On("EmptyHeader")}
}

func (_c *MockProtocol_EmptyHeader_Call) Run(run func()) *MockProtocol_EmptyHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProtocol_EmptyHeader_Call) Return(_a0 networking.Header) *MockProtocol_EmptyHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProtocol_EmptyHeader_Call) RunAndReturn(run func() networking.Header) *MockProtocol_EmptyHeader_Call {
	_c.Call.Return(run)
	return _c
}

// IsAcceptableHandshake provides a mock function with given fields: _a0, _a1
func (_m *MockProtocol) IsAcceptableHandshake(_a0 *networking.Session, _a1 networking.Handshake) bool {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsAcceptableHandshake")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*networking.Session, networking.Handshake) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockProtocol_IsAcceptableHandshake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAcceptableHandshake'
type MockProtocol_IsAcceptableHandshake_Call struct {
	*mock.Call
}

// IsAcceptableHandshake is a helper method to define mock.On call
//   - _a0 *networking.Session
//   - _a1 networking.Handshake
func (_e *MockProtocol_Expecter) IsAcceptableHandshake(_a0 interface{}, _a1 interface{}) *MockProtocol_IsAcceptableHandshake_Call {
	return &MockProtocol_IsAcceptableHandshake_Call{Call: _e.mock.On("IsAcceptableHandshake", _a0, _a1)}
}

func (_c *MockProtocol_IsAcceptableHandshake_Call) Run(run func(_a0 *networking.Session, _a1 networking.Handshake)) *MockProtocol_IsAcceptableHandshake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*networking.Session), args[1].(networking.Handshake))
	})
	return _c
}

func (_c *MockProtocol_IsAcceptableHandshake_Call) Return(_a0 bool) *MockProtocol_IsAcceptableHandshake_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProtocol_IsAcceptableHandshake_Call) RunAndReturn(run func(*networking.Session, networking.Handshake) bool) *MockProtocol_IsAcceptableHandshake_Call {
	_c.Call.Return(run)
	return _c
}

// IsAcceptableMessage provides a mock function with given fields: _a0, _a1
func (_m *MockProtocol) IsAcceptableMessage(_a0 *networking.Session, _a1 networking.Header) bool {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsAcceptableMessage")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*networking.Session, networking.Header) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockProtocol_IsAcceptableMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAcceptableMessage'
type MockProtocol_IsAcceptableMessage_Call struct {
	*mock.Call
}

// IsAcceptableMessage is a helper method to define mock.On call
//   - _a0 *networking.Session
//   - _a1 networking.Header
func (_e *MockProtocol_Expecter) IsAcceptableMessage(_a0 interface{}, _a1 interface{}) *MockProtocol_IsAcceptableMessage_Call {
	return &MockProtocol_IsAcceptableMessage_Call{Call: _e.mock.On("IsAcceptableMessage", _a0, _a1)}
}

func (_c *MockProtocol_IsAcceptableMessage_Call) Run(run func(_a0 *networking.Session, _a1 networking.Header)) *MockProtocol_IsAcceptableMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*networking.Session), args[1].(networking.Header))
	})
	return _c
}

func (_c *MockProtocol_IsAcceptableMessage_Call) Return(_a0 bool) *MockProtocol_IsAcceptableMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProtocol_IsAcceptableMessage_Call) RunAndReturn(run func(*networking.Session, networking.Header) bool) *MockProtocol_IsAcceptableMessage_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with no fields
func (_m *MockProtocol) Ping() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProtocol_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockProtocol_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
func (_e *MockProtocol_Expecter) Ping() *MockProtocol_Ping_Call {
	return &MockProtocol_Ping_Call{Call: _e.mock.On("Ping")}
}

func (_c *MockProtocol_Ping_Call) Run(run func()) *MockProtocol_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProtocol_Ping_Call) Return(_a0 []byte, _a1 error) *MockProtocol_Ping_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProtocol_Ping_Call) RunAndReturn(run func() ([]byte, error)) *MockProtocol_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProtocol creates a new instance of MockProtocol. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProtocol(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProtocol {
	mock := &MockProtocol{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
