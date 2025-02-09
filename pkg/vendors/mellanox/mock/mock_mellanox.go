// Code generated by MockGen. DO NOT EDIT.
// Source: mellanox.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_mellanox.go -source mellanox.go
//

// Package mock_mlxutils is a generated GoMock package.
package mock_mlxutils

import (
	reflect "reflect"

	mlxutils "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/vendors/mellanox"
	gomock "go.uber.org/mock/gomock"
)

// MockMellanoxInterface is a mock of MellanoxInterface interface.
type MockMellanoxInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMellanoxInterfaceMockRecorder
	isgomock struct{}
}

// MockMellanoxInterfaceMockRecorder is the mock recorder for MockMellanoxInterface.
type MockMellanoxInterfaceMockRecorder struct {
	mock *MockMellanoxInterface
}

// NewMockMellanoxInterface creates a new mock instance.
func NewMockMellanoxInterface(ctrl *gomock.Controller) *MockMellanoxInterface {
	mock := &MockMellanoxInterface{ctrl: ctrl}
	mock.recorder = &MockMellanoxInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMellanoxInterface) EXPECT() *MockMellanoxInterfaceMockRecorder {
	return m.recorder
}

// GetMellanoxBlueFieldMode mocks base method.
func (m *MockMellanoxInterface) GetMellanoxBlueFieldMode(arg0 string) (mlxutils.BlueFieldMode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMellanoxBlueFieldMode", arg0)
	ret0, _ := ret[0].(mlxutils.BlueFieldMode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMellanoxBlueFieldMode indicates an expected call of GetMellanoxBlueFieldMode.
func (mr *MockMellanoxInterfaceMockRecorder) GetMellanoxBlueFieldMode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMellanoxBlueFieldMode", reflect.TypeOf((*MockMellanoxInterface)(nil).GetMellanoxBlueFieldMode), arg0)
}

// GetMlxNicFwData mocks base method.
func (m *MockMellanoxInterface) GetMlxNicFwData(pciAddress string) (*mlxutils.MlxNic, *mlxutils.MlxNic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMlxNicFwData", pciAddress)
	ret0, _ := ret[0].(*mlxutils.MlxNic)
	ret1, _ := ret[1].(*mlxutils.MlxNic)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMlxNicFwData indicates an expected call of GetMlxNicFwData.
func (mr *MockMellanoxInterfaceMockRecorder) GetMlxNicFwData(pciAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMlxNicFwData", reflect.TypeOf((*MockMellanoxInterface)(nil).GetMlxNicFwData), pciAddress)
}

// MlxConfigFW mocks base method.
func (m *MockMellanoxInterface) MlxConfigFW(attributesToChange map[string]mlxutils.MlxNic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MlxConfigFW", attributesToChange)
	ret0, _ := ret[0].(error)
	return ret0
}

// MlxConfigFW indicates an expected call of MlxConfigFW.
func (mr *MockMellanoxInterfaceMockRecorder) MlxConfigFW(attributesToChange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MlxConfigFW", reflect.TypeOf((*MockMellanoxInterface)(nil).MlxConfigFW), attributesToChange)
}

// MlxResetFW mocks base method.
func (m *MockMellanoxInterface) MlxResetFW(pciAddresses []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MlxResetFW", pciAddresses)
	ret0, _ := ret[0].(error)
	return ret0
}

// MlxResetFW indicates an expected call of MlxResetFW.
func (mr *MockMellanoxInterfaceMockRecorder) MlxResetFW(pciAddresses any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MlxResetFW", reflect.TypeOf((*MockMellanoxInterface)(nil).MlxResetFW), pciAddresses)
}

// MstConfigReadData mocks base method.
func (m *MockMellanoxInterface) MstConfigReadData(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MstConfigReadData", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MstConfigReadData indicates an expected call of MstConfigReadData.
func (mr *MockMellanoxInterfaceMockRecorder) MstConfigReadData(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MstConfigReadData", reflect.TypeOf((*MockMellanoxInterface)(nil).MstConfigReadData), arg0)
}
