// Code generated by MockGen. DO NOT EDIT.
// Source: uuid.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGenerator is a mock of Generator interface
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockGenerator) Generate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(string)
	return ret0
}

// Generate indicates an expected call of Generate
func (mr *MockGeneratorMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockGenerator)(nil).Generate))
}

// Parse mocks base method
func (m *MockGenerator) Parse(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Parse indicates an expected call of Parse
func (mr *MockGeneratorMockRecorder) Parse(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockGenerator)(nil).Parse), uuid)
}
