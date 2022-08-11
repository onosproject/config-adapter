// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onosproject/config-adapter/pkg/gnmiclient (interfaces: GnmiInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	gnmi "github.com/openconfig/gnmi/proto/gnmi"
	reflect "reflect"
)

// MockGnmiInterface is a mock of GnmiInterface interface
type MockGnmiInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGnmiInterfaceMockRecorder
}

// MockGnmiInterfaceMockRecorder is the mock recorder for MockGnmiInterface
type MockGnmiInterfaceMockRecorder struct {
	mock *MockGnmiInterface
}

// NewMockGnmiInterface creates a new mock instance
func NewMockGnmiInterface(ctrl *gomock.Controller) *MockGnmiInterface {
	mock := &MockGnmiInterface{ctrl: ctrl}
	mock.recorder = &MockGnmiInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGnmiInterface) EXPECT() *MockGnmiInterfaceMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockGnmiInterface) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockGnmiInterfaceMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockGnmiInterface)(nil).Address))
}

// CloseClient mocks base method
func (m *MockGnmiInterface) CloseClient() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseClient")
}

// CloseClient indicates an expected call of CloseClient
func (mr *MockGnmiInterfaceMockRecorder) CloseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseClient", reflect.TypeOf((*MockGnmiInterface)(nil).CloseClient))
}

// Delete mocks base method
func (m *MockGnmiInterface) Delete(arg0 context.Context, arg1 *gnmi.Path, arg2, arg3 string, arg4 []*gnmi.Path) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGnmiInterfaceMockRecorder) Delete(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGnmiInterface)(nil).Delete), arg0, arg1, arg2, arg3, arg4)
}

// GetPath mocks base method
func (m *MockGnmiInterface) GetPath(arg0 context.Context, arg1, arg2, arg3 string) (*gnmi.TypedValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*gnmi.TypedValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPath indicates an expected call of GetPath
func (mr *MockGnmiInterfaceMockRecorder) GetPath(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockGnmiInterface)(nil).GetPath), arg0, arg1, arg2, arg3)
}

// Update mocks base method
func (m *MockGnmiInterface) Update(arg0 context.Context, arg1 *gnmi.Path, arg2, arg3 string, arg4 []*gnmi.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGnmiInterfaceMockRecorder) Update(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGnmiInterface)(nil).Update), arg0, arg1, arg2, arg3, arg4)
}