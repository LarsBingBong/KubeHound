// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/DataDog/KubeHound/pkg/globals/types"
	mock "github.com/stretchr/testify/mock"
)

// ClusterRoleBindingIngestor is an autogenerated mock type for the ClusterRoleBindingIngestor type
type ClusterRoleBindingIngestor struct {
	mock.Mock
}

type ClusterRoleBindingIngestor_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterRoleBindingIngestor) EXPECT() *ClusterRoleBindingIngestor_Expecter {
	return &ClusterRoleBindingIngestor_Expecter{mock: &_m.Mock}
}

// Complete provides a mock function with given fields: _a0
func (_m *ClusterRoleBindingIngestor) Complete(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClusterRoleBindingIngestor_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type ClusterRoleBindingIngestor_Complete_Call struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *ClusterRoleBindingIngestor_Expecter) Complete(_a0 interface{}) *ClusterRoleBindingIngestor_Complete_Call {
	return &ClusterRoleBindingIngestor_Complete_Call{Call: _e.mock.On("Complete", _a0)}
}

func (_c *ClusterRoleBindingIngestor_Complete_Call) Run(run func(_a0 context.Context)) *ClusterRoleBindingIngestor_Complete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ClusterRoleBindingIngestor_Complete_Call) Return(_a0 error) *ClusterRoleBindingIngestor_Complete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterRoleBindingIngestor_Complete_Call) RunAndReturn(run func(context.Context) error) *ClusterRoleBindingIngestor_Complete_Call {
	_c.Call.Return(run)
	return _c
}

// IngestClusterRoleBinding provides a mock function with given fields: _a0, _a1
func (_m *ClusterRoleBindingIngestor) IngestClusterRoleBinding(_a0 context.Context, _a1 types.ClusterRoleBindingType) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ClusterRoleBindingType) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IngestClusterRoleBinding'
type ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call struct {
	*mock.Call
}

// IngestClusterRoleBinding is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.ClusterRoleBindingType
func (_e *ClusterRoleBindingIngestor_Expecter) IngestClusterRoleBinding(_a0 interface{}, _a1 interface{}) *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call {
	return &ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call{Call: _e.mock.On("IngestClusterRoleBinding", _a0, _a1)}
}

func (_c *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call) Run(run func(_a0 context.Context, _a1 types.ClusterRoleBindingType)) *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.ClusterRoleBindingType))
	})
	return _c
}

func (_c *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call) Return(_a0 error) *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call) RunAndReturn(run func(context.Context, types.ClusterRoleBindingType) error) *ClusterRoleBindingIngestor_IngestClusterRoleBinding_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewClusterRoleBindingIngestor interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterRoleBindingIngestor creates a new instance of ClusterRoleBindingIngestor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterRoleBindingIngestor(t mockConstructorTestingTNewClusterRoleBindingIngestor) *ClusterRoleBindingIngestor {
	mock := &ClusterRoleBindingIngestor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
