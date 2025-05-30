// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces_test.go
//
// Generated by this command:
//
//	mockgen -source=interfaces_test.go -destination=mocks/rest_interface.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/api/rbac/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

// MockFakeInterface is a mock of FakeInterface interface.
type MockFakeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFakeInterfaceMockRecorder
	isgomock struct{}
}

// MockFakeInterfaceMockRecorder is the mock recorder for MockFakeInterface.
type MockFakeInterfaceMockRecorder struct {
	mock *MockFakeInterface
}

// NewMockFakeInterface creates a new mock instance.
func NewMockFakeInterface(ctrl *gomock.Controller) *MockFakeInterface {
	mock := &MockFakeInterface{ctrl: ctrl}
	mock.recorder = &MockFakeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFakeInterface) EXPECT() *MockFakeInterfaceMockRecorder {
	return m.recorder
}

// APIVersion mocks base method.
func (m *MockFakeInterface) APIVersion() schema.GroupVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(schema.GroupVersion)
	return ret0
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockFakeInterfaceMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockFakeInterface)(nil).APIVersion))
}

// Delete mocks base method.
func (m *MockFakeInterface) Delete() *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFakeInterfaceMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFakeInterface)(nil).Delete))
}

// Get mocks base method.
func (m *MockFakeInterface) Get() *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockFakeInterfaceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFakeInterface)(nil).Get))
}

// GetRateLimiter mocks base method.
func (m *MockFakeInterface) GetRateLimiter() flowcontrol.RateLimiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateLimiter")
	ret0, _ := ret[0].(flowcontrol.RateLimiter)
	return ret0
}

// GetRateLimiter indicates an expected call of GetRateLimiter.
func (mr *MockFakeInterfaceMockRecorder) GetRateLimiter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateLimiter", reflect.TypeOf((*MockFakeInterface)(nil).GetRateLimiter))
}

// Patch mocks base method.
func (m *MockFakeInterface) Patch(pt types.PatchType) *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", pt)
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockFakeInterfaceMockRecorder) Patch(pt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockFakeInterface)(nil).Patch), pt)
}

// Post mocks base method.
func (m *MockFakeInterface) Post() *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post")
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockFakeInterfaceMockRecorder) Post() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockFakeInterface)(nil).Post))
}

// Put mocks base method.
func (m *MockFakeInterface) Put() *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put")
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockFakeInterfaceMockRecorder) Put() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockFakeInterface)(nil).Put))
}

// Verb mocks base method.
func (m *MockFakeInterface) Verb(verb string) *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verb", verb)
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// Verb indicates an expected call of Verb.
func (mr *MockFakeInterfaceMockRecorder) Verb(verb any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verb", reflect.TypeOf((*MockFakeInterface)(nil).Verb), verb)
}

// MockFakeSubjectNamespacesLister is a mock of FakeSubjectNamespacesLister interface.
type MockFakeSubjectNamespacesLister struct {
	ctrl     *gomock.Controller
	recorder *MockFakeSubjectNamespacesListerMockRecorder
	isgomock struct{}
}

// MockFakeSubjectNamespacesListerMockRecorder is the mock recorder for MockFakeSubjectNamespacesLister.
type MockFakeSubjectNamespacesListerMockRecorder struct {
	mock *MockFakeSubjectNamespacesLister
}

// NewMockFakeSubjectNamespacesLister creates a new mock instance.
func NewMockFakeSubjectNamespacesLister(ctrl *gomock.Controller) *MockFakeSubjectNamespacesLister {
	mock := &MockFakeSubjectNamespacesLister{ctrl: ctrl}
	mock.recorder = &MockFakeSubjectNamespacesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFakeSubjectNamespacesLister) EXPECT() *MockFakeSubjectNamespacesListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockFakeSubjectNamespacesLister) List(subjects ...v10.Subject) []v1.Namespace {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range subjects {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]v1.Namespace)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFakeSubjectNamespacesListerMockRecorder) List(subjects ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFakeSubjectNamespacesLister)(nil).List), subjects...)
}
