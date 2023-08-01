// Code generated by MockGen. DO NOT EDIT.
// Source: tenant.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	management "github.com/auth0/go-auth0/management"
	gomock "github.com/golang/mock/gomock"
)

// MockTenantAPI is a mock of TenantAPI interface.
type MockTenantAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTenantAPIMockRecorder
}

// MockTenantAPIMockRecorder is the mock recorder for MockTenantAPI.
type MockTenantAPIMockRecorder struct {
	mock *MockTenantAPI
}

// NewMockTenantAPI creates a new mock instance.
func NewMockTenantAPI(ctrl *gomock.Controller) *MockTenantAPI {
	mock := &MockTenantAPI{ctrl: ctrl}
	mock.recorder = &MockTenantAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTenantAPI) EXPECT() *MockTenantAPIMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockTenantAPI) Read(ctx context.Context, opts ...management.RequestOption) (*management.Tenant, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*management.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTenantAPIMockRecorder) Read(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTenantAPI)(nil).Read), varargs...)
}
