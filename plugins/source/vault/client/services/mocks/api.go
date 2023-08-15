// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/vault/api"
)

// MockSys is a mock of Sys interface.
type MockSys struct {
	ctrl     *gomock.Controller
	recorder *MockSysMockRecorder
}

// MockSysMockRecorder is the mock recorder for MockSys.
type MockSysMockRecorder struct {
	mock *MockSys
}

// NewMockSys creates a new mock instance.
func NewMockSys(ctrl *gomock.Controller) *MockSys {
	mock := &MockSys{ctrl: ctrl}
	mock.recorder = &MockSysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSys) EXPECT() *MockSysMockRecorder {
	return m.recorder
}

// GetPlugin mocks base method.
func (m *MockSys) GetPlugin(arg0 *api.GetPluginInput) (*api.GetPluginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlugin", arg0)
	ret0, _ := ret[0].(*api.GetPluginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlugin indicates an expected call of GetPlugin.
func (mr *MockSysMockRecorder) GetPlugin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlugin", reflect.TypeOf((*MockSys)(nil).GetPlugin), arg0)
}

// GetPluginWithContext mocks base method.
func (m *MockSys) GetPluginWithContext(arg0 context.Context, arg1 *api.GetPluginInput) (*api.GetPluginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginWithContext", arg0, arg1)
	ret0, _ := ret[0].(*api.GetPluginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginWithContext indicates an expected call of GetPluginWithContext.
func (mr *MockSysMockRecorder) GetPluginWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginWithContext", reflect.TypeOf((*MockSys)(nil).GetPluginWithContext), arg0, arg1)
}

// GetPolicy mocks base method.
func (m *MockSys) GetPolicy(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicy", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockSysMockRecorder) GetPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockSys)(nil).GetPolicy), arg0)
}

// GetPolicyWithContext mocks base method.
func (m *MockSys) GetPolicyWithContext(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyWithContext", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyWithContext indicates an expected call of GetPolicyWithContext.
func (mr *MockSysMockRecorder) GetPolicyWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyWithContext", reflect.TypeOf((*MockSys)(nil).GetPolicyWithContext), arg0, arg1)
}

// ListAudit mocks base method.
func (m *MockSys) ListAudit() (map[string]*api.Audit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAudit")
	ret0, _ := ret[0].(map[string]*api.Audit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAudit indicates an expected call of ListAudit.
func (mr *MockSysMockRecorder) ListAudit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAudit", reflect.TypeOf((*MockSys)(nil).ListAudit))
}

// ListAuditWithContext mocks base method.
func (m *MockSys) ListAuditWithContext(arg0 context.Context) (map[string]*api.Audit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuditWithContext", arg0)
	ret0, _ := ret[0].(map[string]*api.Audit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuditWithContext indicates an expected call of ListAuditWithContext.
func (mr *MockSysMockRecorder) ListAuditWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuditWithContext", reflect.TypeOf((*MockSys)(nil).ListAuditWithContext), arg0)
}

// ListAuth mocks base method.
func (m *MockSys) ListAuth() (map[string]*api.MountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuth")
	ret0, _ := ret[0].(map[string]*api.MountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuth indicates an expected call of ListAuth.
func (mr *MockSysMockRecorder) ListAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuth", reflect.TypeOf((*MockSys)(nil).ListAuth))
}

// ListAuthWithContext mocks base method.
func (m *MockSys) ListAuthWithContext(arg0 context.Context) (map[string]*api.MountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthWithContext", arg0)
	ret0, _ := ret[0].(map[string]*api.MountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthWithContext indicates an expected call of ListAuthWithContext.
func (mr *MockSysMockRecorder) ListAuthWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthWithContext", reflect.TypeOf((*MockSys)(nil).ListAuthWithContext), arg0)
}

// ListMounts mocks base method.
func (m *MockSys) ListMounts() (map[string]*api.MountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMounts")
	ret0, _ := ret[0].(map[string]*api.MountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMounts indicates an expected call of ListMounts.
func (mr *MockSysMockRecorder) ListMounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMounts", reflect.TypeOf((*MockSys)(nil).ListMounts))
}

// ListMountsWithContext mocks base method.
func (m *MockSys) ListMountsWithContext(arg0 context.Context) (map[string]*api.MountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMountsWithContext", arg0)
	ret0, _ := ret[0].(map[string]*api.MountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMountsWithContext indicates an expected call of ListMountsWithContext.
func (mr *MockSysMockRecorder) ListMountsWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMountsWithContext", reflect.TypeOf((*MockSys)(nil).ListMountsWithContext), arg0)
}

// ListPlugins mocks base method.
func (m *MockSys) ListPlugins(arg0 *api.ListPluginsInput) (*api.ListPluginsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlugins", arg0)
	ret0, _ := ret[0].(*api.ListPluginsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlugins indicates an expected call of ListPlugins.
func (mr *MockSysMockRecorder) ListPlugins(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlugins", reflect.TypeOf((*MockSys)(nil).ListPlugins), arg0)
}

// ListPluginsWithContext mocks base method.
func (m *MockSys) ListPluginsWithContext(arg0 context.Context, arg1 *api.ListPluginsInput) (*api.ListPluginsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPluginsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*api.ListPluginsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPluginsWithContext indicates an expected call of ListPluginsWithContext.
func (mr *MockSysMockRecorder) ListPluginsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPluginsWithContext", reflect.TypeOf((*MockSys)(nil).ListPluginsWithContext), arg0, arg1)
}

// ListPolicies mocks base method.
func (m *MockSys) ListPolicies() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolicies")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockSysMockRecorder) ListPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockSys)(nil).ListPolicies))
}

// ListPoliciesWithContext mocks base method.
func (m *MockSys) ListPoliciesWithContext(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPoliciesWithContext", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPoliciesWithContext indicates an expected call of ListPoliciesWithContext.
func (mr *MockSysMockRecorder) ListPoliciesWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPoliciesWithContext", reflect.TypeOf((*MockSys)(nil).ListPoliciesWithContext), arg0)
}