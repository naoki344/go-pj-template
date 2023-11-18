// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/naoki.miyoshi/Data/pjct/en/en/backend/test/scripts/../../internal/adapter/driver/ogenlib/oas_server_gen.go

// Package mock_ogen is a generated GoMock package.
package mock_ogen

import (
	context "context"
	reflect "reflect"

	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetCustomerByID mocks base method.
func (m *MockHandler) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (ogen.GetCustomerByIDRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerByID", ctx, params)
	ret0, _ := ret[0].(ogen.GetCustomerByIDRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerByID indicates an expected call of GetCustomerByID.
func (mr *MockHandlerMockRecorder) GetCustomerByID(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerByID", reflect.TypeOf((*MockHandler)(nil).GetCustomerByID), ctx, params)
}

// PostCreateCustomer mocks base method.
func (m *MockHandler) PostCreateCustomer(ctx context.Context, req *ogen.PostCreateCustomerRequest) (ogen.PostCreateCustomerRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCreateCustomer", ctx, req)
	ret0, _ := ret[0].(ogen.PostCreateCustomerRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCreateCustomer indicates an expected call of PostCreateCustomer.
func (mr *MockHandlerMockRecorder) PostCreateCustomer(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCreateCustomer", reflect.TypeOf((*MockHandler)(nil).PostCreateCustomer), ctx, req)
}

// PostSearchCustomer mocks base method.
func (m *MockHandler) PostSearchCustomer(ctx context.Context, req *ogen.PostSearchCustomerRequest) (ogen.PostSearchCustomerRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostSearchCustomer", ctx, req)
	ret0, _ := ret[0].(ogen.PostSearchCustomerRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSearchCustomer indicates an expected call of PostSearchCustomer.
func (mr *MockHandlerMockRecorder) PostSearchCustomer(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSearchCustomer", reflect.TypeOf((*MockHandler)(nil).PostSearchCustomer), ctx, req)
}

// PutModifyCustomerByID mocks base method.
func (m *MockHandler) PutModifyCustomerByID(ctx context.Context, req *ogen.PutModifyCustomerByIDRequest, params ogen.PutModifyCustomerByIDParams) (ogen.PutModifyCustomerByIDRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutModifyCustomerByID", ctx, req, params)
	ret0, _ := ret[0].(ogen.PutModifyCustomerByIDRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutModifyCustomerByID indicates an expected call of PutModifyCustomerByID.
func (mr *MockHandlerMockRecorder) PutModifyCustomerByID(ctx, req, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutModifyCustomerByID", reflect.TypeOf((*MockHandler)(nil).PutModifyCustomerByID), ctx, req, params)
}
