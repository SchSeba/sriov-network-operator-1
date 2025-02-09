// Code generated by MockGen. DO NOT EDIT.
// Source: utils.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_utils.go -source utils.go
//

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCmdInterface is a mock of CmdInterface interface.
type MockCmdInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCmdInterfaceMockRecorder
	isgomock struct{}
}

// MockCmdInterfaceMockRecorder is the mock recorder for MockCmdInterface.
type MockCmdInterfaceMockRecorder struct {
	mock *MockCmdInterface
}

// NewMockCmdInterface creates a new mock instance.
func NewMockCmdInterface(ctrl *gomock.Controller) *MockCmdInterface {
	mock := &MockCmdInterface{ctrl: ctrl}
	mock.recorder = &MockCmdInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmdInterface) EXPECT() *MockCmdInterfaceMockRecorder {
	return m.recorder
}

// Chroot mocks base method.
func (m *MockCmdInterface) Chroot(arg0 string) (func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chroot", arg0)
	ret0, _ := ret[0].(func() error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chroot indicates an expected call of Chroot.
func (mr *MockCmdInterfaceMockRecorder) Chroot(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chroot", reflect.TypeOf((*MockCmdInterface)(nil).Chroot), arg0)
}

// RunCommand mocks base method.
func (m *MockCmdInterface) RunCommand(arg0 string, arg1 ...string) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommand", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockCmdInterfaceMockRecorder) RunCommand(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockCmdInterface)(nil).RunCommand), varargs...)
}
