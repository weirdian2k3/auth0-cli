// Code generated by MockGen. DO NOT EDIT.
// Source: resource_server.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	management "github.com/auth0/go-auth0/management"
	gomock "github.com/golang/mock/gomock"
)

// MockResourceServerAPI is a mock of ResourceServerAPI interface.
type MockResourceServerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockResourceServerAPIMockRecorder
}

// MockResourceServerAPIMockRecorder is the mock recorder for MockResourceServerAPI.
type MockResourceServerAPIMockRecorder struct {
	mock *MockResourceServerAPI
}

// NewMockResourceServerAPI creates a new mock instance.
func NewMockResourceServerAPI(ctrl *gomock.Controller) *MockResourceServerAPI {
	mock := &MockResourceServerAPI{ctrl: ctrl}
	mock.recorder = &MockResourceServerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceServerAPI) EXPECT() *MockResourceServerAPIMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockResourceServerAPI) Create(ctx context.Context, rs *management.ResourceServer, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResourceServerAPIMockRecorder) Create(ctx, rs interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceServerAPI)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockResourceServerAPI) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResourceServerAPIMockRecorder) Delete(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceServerAPI)(nil).Delete), varargs...)
}

// List mocks base method.
func (m *MockResourceServerAPI) List(ctx context.Context, opts ...management.RequestOption) (*management.ResourceServerList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*management.ResourceServerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourceServerAPIMockRecorder) List(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceServerAPI)(nil).List), varargs...)
}

// Read mocks base method.
func (m *MockResourceServerAPI) Read(ctx context.Context, id string, opts ...management.RequestOption) (*management.ResourceServer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*management.ResourceServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockResourceServerAPIMockRecorder) Read(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockResourceServerAPI)(nil).Read), varargs...)
}

// Update mocks base method.
func (m *MockResourceServerAPI) Update(ctx context.Context, id string, rs *management.ResourceServer, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, rs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResourceServerAPIMockRecorder) Update(ctx, id, rs interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, rs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceServerAPI)(nil).Update), varargs...)
}
