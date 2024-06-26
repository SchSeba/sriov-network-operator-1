// Code generated by MockGen. DO NOT EDIT.
// Source: ghw.go

// Package mock_ghw is a generated GoMock package.
package mock_ghw

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ghw "github.com/jaypipes/ghw"
	ghw0 "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/host/internal/lib/ghw"
)

// MockGHWLib is a mock of GHWLib interface.
type MockGHWLib struct {
	ctrl     *gomock.Controller
	recorder *MockGHWLibMockRecorder
}

// MockGHWLibMockRecorder is the mock recorder for MockGHWLib.
type MockGHWLibMockRecorder struct {
	mock *MockGHWLib
}

// NewMockGHWLib creates a new mock instance.
func NewMockGHWLib(ctrl *gomock.Controller) *MockGHWLib {
	mock := &MockGHWLib{ctrl: ctrl}
	mock.recorder = &MockGHWLibMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGHWLib) EXPECT() *MockGHWLibMockRecorder {
	return m.recorder
}

// PCI mocks base method.
func (m *MockGHWLib) PCI() (ghw0.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PCI")
	ret0, _ := ret[0].(ghw0.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PCI indicates an expected call of PCI.
func (mr *MockGHWLibMockRecorder) PCI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PCI", reflect.TypeOf((*MockGHWLib)(nil).PCI))
}

// MockInfo is a mock of Info interface.
type MockInfo struct {
	ctrl     *gomock.Controller
	recorder *MockInfoMockRecorder
}

// MockInfoMockRecorder is the mock recorder for MockInfo.
type MockInfoMockRecorder struct {
	mock *MockInfo
}

// NewMockInfo creates a new mock instance.
func NewMockInfo(ctrl *gomock.Controller) *MockInfo {
	mock := &MockInfo{ctrl: ctrl}
	mock.recorder = &MockInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfo) EXPECT() *MockInfoMockRecorder {
	return m.recorder
}

// ListDevices mocks base method.
func (m *MockInfo) ListDevices() []*ghw.PCIDevice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDevices")
	ret0, _ := ret[0].([]*ghw.PCIDevice)
	return ret0
}

// ListDevices indicates an expected call of ListDevices.
func (mr *MockInfoMockRecorder) ListDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevices", reflect.TypeOf((*MockInfo)(nil).ListDevices))
}
