// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dubbogo/arana/pkg/proto/rule (interfaces: ShardComputer)

// Package testdata is a generated GoMock package.
package testdata

import (
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

// MockShardComputer is a mock of ShardComputer interface.
type MockShardComputer struct {
	ctrl     *gomock.Controller
	recorder *MockShardComputerMockRecorder
}

// MockShardComputerMockRecorder is the mock recorder for MockShardComputer.
type MockShardComputerMockRecorder struct {
	mock *MockShardComputer
}

// NewMockShardComputer creates a new mock instance.
func NewMockShardComputer(ctrl *gomock.Controller) *MockShardComputer {
	mock := &MockShardComputer{ctrl: ctrl}
	mock.recorder = &MockShardComputerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShardComputer) EXPECT() *MockShardComputerMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockShardComputer) Compute(arg0 interface{}) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compute", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compute indicates an expected call of Compute.
func (mr *MockShardComputerMockRecorder) Compute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockShardComputer)(nil).Compute), arg0)
}
