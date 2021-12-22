// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dubbogo/arana/pkg/proto (interfaces: Rows,VConn,QueryPlan,ExecPlan,Optimizer)

// Package testdata is a generated GoMock package.
package testdata

import (
	context "context"
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

import (
	proto "github.com/dubbogo/arana/pkg/proto"
)

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockRows) Next() proto.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(proto.Row)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowsMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRows)(nil).Next))
}

// MockVConn is a mock of VConn interface.
type MockVConn struct {
	ctrl     *gomock.Controller
	recorder *MockVConnMockRecorder
}

// MockVConnMockRecorder is the mock recorder for MockVConn.
type MockVConnMockRecorder struct {
	mock *MockVConn
}

// NewMockVConn creates a new mock instance.
func NewMockVConn(ctrl *gomock.Controller) *MockVConn {
	mock := &MockVConn{ctrl: ctrl}
	mock.recorder = &MockVConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVConn) EXPECT() *MockVConnMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockVConn) Exec(arg0 context.Context, arg1, arg2 string, arg3 ...interface{}) (proto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockVConnMockRecorder) Exec(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockVConn)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockVConn) Query(arg0 context.Context, arg1, arg2 string, arg3 ...interface{}) (proto.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(proto.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockVConnMockRecorder) Query(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockVConn)(nil).Query), varargs...)
}

// MockQueryPlan is a mock of QueryPlan interface.
type MockQueryPlan struct {
	ctrl     *gomock.Controller
	recorder *MockQueryPlanMockRecorder
}

// MockQueryPlanMockRecorder is the mock recorder for MockQueryPlan.
type MockQueryPlanMockRecorder struct {
	mock *MockQueryPlan
}

// NewMockQueryPlan creates a new mock instance.
func NewMockQueryPlan(ctrl *gomock.Controller) *MockQueryPlan {
	mock := &MockQueryPlan{ctrl: ctrl}
	mock.recorder = &MockQueryPlanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryPlan) EXPECT() *MockQueryPlanMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockQueryPlan) Query(arg0 context.Context, arg1 *proto.VConn) (proto.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(proto.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockQueryPlanMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockQueryPlan)(nil).Query), arg0, arg1)
}

// MockExecPlan is a mock of ExecPlan interface.
type MockExecPlan struct {
	ctrl     *gomock.Controller
	recorder *MockExecPlanMockRecorder
}

// MockExecPlanMockRecorder is the mock recorder for MockExecPlan.
type MockExecPlanMockRecorder struct {
	mock *MockExecPlan
}

// NewMockExecPlan creates a new mock instance.
func NewMockExecPlan(ctrl *gomock.Controller) *MockExecPlan {
	mock := &MockExecPlan{ctrl: ctrl}
	mock.recorder = &MockExecPlanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecPlan) EXPECT() *MockExecPlanMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockExecPlan) Exec(arg0 context.Context, arg1 *proto.VConn) (proto.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(proto.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockExecPlanMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExecPlan)(nil).Exec), arg0, arg1)
}

// MockOptimizer is a mock of Optimizer interface.
type MockOptimizer struct {
	ctrl     *gomock.Controller
	recorder *MockOptimizerMockRecorder
}

// MockOptimizerMockRecorder is the mock recorder for MockOptimizer.
type MockOptimizerMockRecorder struct {
	mock *MockOptimizer
}

// NewMockOptimizer creates a new mock instance.
func NewMockOptimizer(ctrl *gomock.Controller) *MockOptimizer {
	mock := &MockOptimizer{ctrl: ctrl}
	mock.recorder = &MockOptimizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOptimizer) EXPECT() *MockOptimizerMockRecorder {
	return m.recorder
}

// Optimize mocks base method.
func (m *MockOptimizer) Optimize(arg0 context.Context, arg1 string, arg2 ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Optimize", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Optimize indicates an expected call of Optimize.
func (mr *MockOptimizerMockRecorder) Optimize(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Optimize", reflect.TypeOf((*MockOptimizer)(nil).Optimize), varargs...)
}
