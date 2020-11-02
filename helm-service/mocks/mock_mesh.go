// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn/keptn/helm-service/pkg/mesh (interfaces: Mesh)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mesh "github.com/keptn/keptn/helm-service/pkg/mesh"
)

// MockMesh is a mock of Mesh interface.
type MockMesh struct {
	ctrl     *gomock.Controller
	recorder *MockMeshMockRecorder
}

// MockMeshMockRecorder is the mock recorder for MockMesh.
type MockMeshMockRecorder struct {
	mock *MockMesh
}

// NewMockMesh creates a new mock instance.
func NewMockMesh(ctrl *gomock.Controller) *MockMesh {
	mock := &MockMesh{ctrl: ctrl}
	mock.recorder = &MockMeshMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMesh) EXPECT() *MockMeshMockRecorder {
	return m.recorder
}

// GenerateDestinationRule mocks base method.
func (m *MockMesh) GenerateDestinationRule(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDestinationRule", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateDestinationRule indicates an expected call of GenerateDestinationRule.
func (mr *MockMeshMockRecorder) GenerateDestinationRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDestinationRule", reflect.TypeOf((*MockMesh)(nil).GenerateDestinationRule), arg0, arg1)
}

// GenerateVirtualService mocks base method.
func (m *MockMesh) GenerateVirtualService(arg0 string, arg1, arg2 []string, arg3 []mesh.HTTPRouteDestination) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateVirtualService", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateVirtualService indicates an expected call of GenerateVirtualService.
func (mr *MockMeshMockRecorder) GenerateVirtualService(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateVirtualService", reflect.TypeOf((*MockMesh)(nil).GenerateVirtualService), arg0, arg1, arg2, arg3)
}

// GetDestinationRuleSuffix mocks base method.
func (m *MockMesh) GetDestinationRuleSuffix() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDestinationRuleSuffix")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDestinationRuleSuffix indicates an expected call of GetDestinationRuleSuffix.
func (mr *MockMeshMockRecorder) GetDestinationRuleSuffix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDestinationRuleSuffix", reflect.TypeOf((*MockMesh)(nil).GetDestinationRuleSuffix))
}

// GetVirtualServiceSuffix mocks base method.
func (m *MockMesh) GetVirtualServiceSuffix() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualServiceSuffix")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVirtualServiceSuffix indicates an expected call of GetVirtualServiceSuffix.
func (mr *MockMeshMockRecorder) GetVirtualServiceSuffix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualServiceSuffix", reflect.TypeOf((*MockMesh)(nil).GetVirtualServiceSuffix))
}

// UpdateWeights mocks base method.
func (m *MockMesh) UpdateWeights(arg0 []byte, arg1 int32) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWeights", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWeights indicates an expected call of UpdateWeights.
func (mr *MockMeshMockRecorder) UpdateWeights(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWeights", reflect.TypeOf((*MockMesh)(nil).UpdateWeights), arg0, arg1)
}