// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/naoki.miyoshi/Data/pjct/en/en/backend/test/scripts/../../internal/adapter/driver/lambda/interface.go

// Package mock_lambdaadapter is a generated GoMock package.
package mock_lambdaadapter

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLambdaHandler is a mock of LambdaHandler interface.
type MockLambdaHandler struct {
	ctrl     *gomock.Controller
	recorder *MockLambdaHandlerMockRecorder
}

// MockLambdaHandlerMockRecorder is the mock recorder for MockLambdaHandler.
type MockLambdaHandlerMockRecorder struct {
	mock *MockLambdaHandler
}

// NewMockLambdaHandler creates a new mock instance.
func NewMockLambdaHandler(ctrl *gomock.Controller) *MockLambdaHandler {
	mock := &MockLambdaHandler{ctrl: ctrl}
	mock.recorder = &MockLambdaHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLambdaHandler) EXPECT() *MockLambdaHandlerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockLambdaHandler) Run() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockLambdaHandlerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockLambdaHandler)(nil).Run))
}