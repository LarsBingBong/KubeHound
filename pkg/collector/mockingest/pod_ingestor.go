// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/DataDog/KubeHound/pkg/globals/types"
	mock "github.com/stretchr/testify/mock"
)

// PodIngestor is an autogenerated mock type for the PodIngestor type
type PodIngestor struct {
	mock.Mock
}

type PodIngestor_Expecter struct {
	mock *mock.Mock
}

func (_m *PodIngestor) EXPECT() *PodIngestor_Expecter {
	return &PodIngestor_Expecter{mock: &_m.Mock}
}

// Complete provides a mock function with given fields: _a0
func (_m *PodIngestor) Complete(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodIngestor_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type PodIngestor_Complete_Call struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PodIngestor_Expecter) Complete(_a0 interface{}) *PodIngestor_Complete_Call {
	return &PodIngestor_Complete_Call{Call: _e.mock.On("Complete", _a0)}
}

func (_c *PodIngestor_Complete_Call) Run(run func(_a0 context.Context)) *PodIngestor_Complete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PodIngestor_Complete_Call) Return(_a0 error) *PodIngestor_Complete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodIngestor_Complete_Call) RunAndReturn(run func(context.Context) error) *PodIngestor_Complete_Call {
	_c.Call.Return(run)
	return _c
}

// IngestPod provides a mock function with given fields: _a0, _a1
func (_m *PodIngestor) IngestPod(_a0 context.Context, _a1 types.PodType) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.PodType) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodIngestor_IngestPod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IngestPod'
type PodIngestor_IngestPod_Call struct {
	*mock.Call
}

// IngestPod is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.PodType
func (_e *PodIngestor_Expecter) IngestPod(_a0 interface{}, _a1 interface{}) *PodIngestor_IngestPod_Call {
	return &PodIngestor_IngestPod_Call{Call: _e.mock.On("IngestPod", _a0, _a1)}
}

func (_c *PodIngestor_IngestPod_Call) Run(run func(_a0 context.Context, _a1 types.PodType)) *PodIngestor_IngestPod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.PodType))
	})
	return _c
}

func (_c *PodIngestor_IngestPod_Call) Return(_a0 error) *PodIngestor_IngestPod_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodIngestor_IngestPod_Call) RunAndReturn(run func(context.Context, types.PodType) error) *PodIngestor_IngestPod_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPodIngestor interface {
	mock.TestingT
	Cleanup(func())
}

// NewPodIngestor creates a new instance of PodIngestor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodIngestor(t mockConstructorTestingTNewPodIngestor) *PodIngestor {
	mock := &PodIngestor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
