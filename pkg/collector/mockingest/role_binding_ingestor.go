// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/DataDog/KubeHound/pkg/globals/types"
	mock "github.com/stretchr/testify/mock"
)

// RoleBindingIngestor is an autogenerated mock type for the RoleBindingIngestor type
type RoleBindingIngestor struct {
	mock.Mock
}

type RoleBindingIngestor_Expecter struct {
	mock *mock.Mock
}

func (_m *RoleBindingIngestor) EXPECT() *RoleBindingIngestor_Expecter {
	return &RoleBindingIngestor_Expecter{mock: &_m.Mock}
}

// Complete provides a mock function with given fields: _a0
func (_m *RoleBindingIngestor) Complete(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleBindingIngestor_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type RoleBindingIngestor_Complete_Call struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *RoleBindingIngestor_Expecter) Complete(_a0 interface{}) *RoleBindingIngestor_Complete_Call {
	return &RoleBindingIngestor_Complete_Call{Call: _e.mock.On("Complete", _a0)}
}

func (_c *RoleBindingIngestor_Complete_Call) Run(run func(_a0 context.Context)) *RoleBindingIngestor_Complete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RoleBindingIngestor_Complete_Call) Return(_a0 error) *RoleBindingIngestor_Complete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoleBindingIngestor_Complete_Call) RunAndReturn(run func(context.Context) error) *RoleBindingIngestor_Complete_Call {
	_c.Call.Return(run)
	return _c
}

// IngestRoleBinding provides a mock function with given fields: _a0, _a1
func (_m *RoleBindingIngestor) IngestRoleBinding(_a0 context.Context, _a1 types.RoleBindingType) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RoleBindingType) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleBindingIngestor_IngestRoleBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IngestRoleBinding'
type RoleBindingIngestor_IngestRoleBinding_Call struct {
	*mock.Call
}

// IngestRoleBinding is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.RoleBindingType
func (_e *RoleBindingIngestor_Expecter) IngestRoleBinding(_a0 interface{}, _a1 interface{}) *RoleBindingIngestor_IngestRoleBinding_Call {
	return &RoleBindingIngestor_IngestRoleBinding_Call{Call: _e.mock.On("IngestRoleBinding", _a0, _a1)}
}

func (_c *RoleBindingIngestor_IngestRoleBinding_Call) Run(run func(_a0 context.Context, _a1 types.RoleBindingType)) *RoleBindingIngestor_IngestRoleBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.RoleBindingType))
	})
	return _c
}

func (_c *RoleBindingIngestor_IngestRoleBinding_Call) Return(_a0 error) *RoleBindingIngestor_IngestRoleBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoleBindingIngestor_IngestRoleBinding_Call) RunAndReturn(run func(context.Context, types.RoleBindingType) error) *RoleBindingIngestor_IngestRoleBinding_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRoleBindingIngestor interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleBindingIngestor creates a new instance of RoleBindingIngestor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleBindingIngestor(t mockConstructorTestingTNewRoleBindingIngestor) *RoleBindingIngestor {
	mock := &RoleBindingIngestor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
