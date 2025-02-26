// Code generated by MockGen. DO NOT EDIT.
// Source: dputils.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_dputils.go -source dputils.go
//

// Package mock_dputils is a generated GoMock package.
package mock_dputils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDPUtilsLib is a mock of DPUtilsLib interface.
type MockDPUtilsLib struct {
	ctrl     *gomock.Controller
	recorder *MockDPUtilsLibMockRecorder
	isgomock struct{}
}

// MockDPUtilsLibMockRecorder is the mock recorder for MockDPUtilsLib.
type MockDPUtilsLibMockRecorder struct {
	mock *MockDPUtilsLib
}

// NewMockDPUtilsLib creates a new mock instance.
func NewMockDPUtilsLib(ctrl *gomock.Controller) *MockDPUtilsLib {
	mock := &MockDPUtilsLib{ctrl: ctrl}
	mock.recorder = &MockDPUtilsLibMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDPUtilsLib) EXPECT() *MockDPUtilsLibMockRecorder {
	return m.recorder
}

// GetDriverName mocks base method.
func (m *MockDPUtilsLib) GetDriverName(pciAddr string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverName", pciAddr)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverName indicates an expected call of GetDriverName.
func (mr *MockDPUtilsLibMockRecorder) GetDriverName(pciAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverName", reflect.TypeOf((*MockDPUtilsLib)(nil).GetDriverName), pciAddr)
}

// GetNetNames mocks base method.
func (m *MockDPUtilsLib) GetNetNames(pciAddr string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetNames", pciAddr)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetNames indicates an expected call of GetNetNames.
func (mr *MockDPUtilsLibMockRecorder) GetNetNames(pciAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetNames", reflect.TypeOf((*MockDPUtilsLib)(nil).GetNetNames), pciAddr)
}

// GetSriovVFcapacity mocks base method.
func (m *MockDPUtilsLib) GetSriovVFcapacity(pf string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSriovVFcapacity", pf)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSriovVFcapacity indicates an expected call of GetSriovVFcapacity.
func (mr *MockDPUtilsLibMockRecorder) GetSriovVFcapacity(pf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSriovVFcapacity", reflect.TypeOf((*MockDPUtilsLib)(nil).GetSriovVFcapacity), pf)
}

// GetVFID mocks base method.
func (m *MockDPUtilsLib) GetVFID(pciAddr string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVFID", pciAddr)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVFID indicates an expected call of GetVFID.
func (mr *MockDPUtilsLibMockRecorder) GetVFID(pciAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVFID", reflect.TypeOf((*MockDPUtilsLib)(nil).GetVFID), pciAddr)
}

// GetVFList mocks base method.
func (m *MockDPUtilsLib) GetVFList(pf string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVFList", pf)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVFList indicates an expected call of GetVFList.
func (mr *MockDPUtilsLibMockRecorder) GetVFList(pf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVFList", reflect.TypeOf((*MockDPUtilsLib)(nil).GetVFList), pf)
}

// GetVFconfigured mocks base method.
func (m *MockDPUtilsLib) GetVFconfigured(pf string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVFconfigured", pf)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetVFconfigured indicates an expected call of GetVFconfigured.
func (mr *MockDPUtilsLibMockRecorder) GetVFconfigured(pf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVFconfigured", reflect.TypeOf((*MockDPUtilsLib)(nil).GetVFconfigured), pf)
}

// IsSriovPF mocks base method.
func (m *MockDPUtilsLib) IsSriovPF(pciAddr string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSriovPF", pciAddr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSriovPF indicates an expected call of IsSriovPF.
func (mr *MockDPUtilsLibMockRecorder) IsSriovPF(pciAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSriovPF", reflect.TypeOf((*MockDPUtilsLib)(nil).IsSriovPF), pciAddr)
}

// IsSriovVF mocks base method.
func (m *MockDPUtilsLib) IsSriovVF(pciAddr string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSriovVF", pciAddr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSriovVF indicates an expected call of IsSriovVF.
func (mr *MockDPUtilsLibMockRecorder) IsSriovVF(pciAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSriovVF", reflect.TypeOf((*MockDPUtilsLib)(nil).IsSriovVF), pciAddr)
}

// SriovConfigured mocks base method.
func (m *MockDPUtilsLib) SriovConfigured(addr string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SriovConfigured", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SriovConfigured indicates an expected call of SriovConfigured.
func (mr *MockDPUtilsLibMockRecorder) SriovConfigured(addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SriovConfigured", reflect.TypeOf((*MockDPUtilsLib)(nil).SriovConfigured), addr)
}
