// Code generated by MockGen. DO NOT EDIT.
// Source: organization.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	management "github.com/auth0/go-auth0/management"
	gomock "github.com/golang/mock/gomock"
)

// MockOrganizationAPI is a mock of OrganizationAPI interface.
type MockOrganizationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIMockRecorder
}

// MockOrganizationAPIMockRecorder is the mock recorder for MockOrganizationAPI.
type MockOrganizationAPIMockRecorder struct {
	mock *MockOrganizationAPI
}

// NewMockOrganizationAPI creates a new mock instance.
func NewMockOrganizationAPI(ctrl *gomock.Controller) *MockOrganizationAPI {
	mock := &MockOrganizationAPI{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPI) EXPECT() *MockOrganizationAPIMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganizationAPI) Create(ctx context.Context, o *management.Organization, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, o}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationAPIMockRecorder) Create(ctx, o interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, o}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationAPI)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockOrganizationAPI) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {
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
func (mr *MockOrganizationAPIMockRecorder) Delete(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationAPI)(nil).Delete), varargs...)
}

// List mocks base method.
func (m *MockOrganizationAPI) List(ctx context.Context, opts ...management.RequestOption) (*management.OrganizationList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*management.OrganizationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrganizationAPIMockRecorder) List(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationAPI)(nil).List), varargs...)
}

// MemberRoles mocks base method.
func (m *MockOrganizationAPI) MemberRoles(ctx context.Context, id, userID string, opts ...management.RequestOption) (*management.OrganizationMemberRoleList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, userID}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MemberRoles", varargs...)
	ret0, _ := ret[0].(*management.OrganizationMemberRoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberRoles indicates an expected call of MemberRoles.
func (mr *MockOrganizationAPIMockRecorder) MemberRoles(ctx, id, userID interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, userID}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberRoles", reflect.TypeOf((*MockOrganizationAPI)(nil).MemberRoles), varargs...)
}

// Members mocks base method.
func (m *MockOrganizationAPI) Members(ctx context.Context, id string, opts ...management.RequestOption) (*management.OrganizationMemberList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Members", varargs...)
	ret0, _ := ret[0].(*management.OrganizationMemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Members indicates an expected call of Members.
func (mr *MockOrganizationAPIMockRecorder) Members(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockOrganizationAPI)(nil).Members), varargs...)
}

// Read mocks base method.
func (m *MockOrganizationAPI) Read(ctx context.Context, id string, opts ...management.RequestOption) (*management.Organization, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*management.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOrganizationAPIMockRecorder) Read(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOrganizationAPI)(nil).Read), varargs...)
}

// Update mocks base method.
func (m *MockOrganizationAPI) Update(ctx context.Context, id string, o *management.Organization, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, o}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationAPIMockRecorder) Update(ctx, id, o interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, o}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationAPI)(nil).Update), varargs...)
}
