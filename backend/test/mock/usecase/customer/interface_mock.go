// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/naoki.miyoshi/Data/pjct/en/en/backend/test/scripts/../../internal/usecase/customer/interface.go

// Package mock_customerusecase is a generated GoMock package.
package mock_customerusecase

import (
	reflect "reflect"

	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomerUsecaseInterface is a mock of CustomerUsecaseInterface interface.
type MockCustomerUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerUsecaseInterfaceMockRecorder
}

// MockCustomerUsecaseInterfaceMockRecorder is the mock recorder for MockCustomerUsecaseInterface.
type MockCustomerUsecaseInterfaceMockRecorder struct {
	mock *MockCustomerUsecaseInterface
}

// NewMockCustomerUsecaseInterface creates a new mock instance.
func NewMockCustomerUsecaseInterface(ctrl *gomock.Controller) *MockCustomerUsecaseInterface {
	mock := &MockCustomerUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockCustomerUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerUsecaseInterface) EXPECT() *MockCustomerUsecaseInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerUsecaseInterface) Create(customer *customermodel.Customer) (*customermodel.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", customer)
	ret0, _ := ret[0].(*customermodel.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerUsecaseInterfaceMockRecorder) Create(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerUsecaseInterface)(nil).Create), customer)
}

// GetByID mocks base method.
func (m *MockCustomerUsecaseInterface) GetByID(customerID customermodel.ID) (*customermodel.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", customerID)
	ret0, _ := ret[0].(*customermodel.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCustomerUsecaseInterfaceMockRecorder) GetByID(customerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCustomerUsecaseInterface)(nil).GetByID), customerID)
}

// Search mocks base method.
func (m *MockCustomerUsecaseInterface) Search(pageNumber, pageSize int64, conditions *customermodel.SearchConditions) (*[]*customermodel.Customer, *pagemodel.PageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", pageNumber, pageSize, conditions)
	ret0, _ := ret[0].(*[]*customermodel.Customer)
	ret1, _ := ret[1].(*pagemodel.PageResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Search indicates an expected call of Search.
func (mr *MockCustomerUsecaseInterfaceMockRecorder) Search(pageNumber, pageSize, conditions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockCustomerUsecaseInterface)(nil).Search), pageNumber, pageSize, conditions)
}

// UpdateByID mocks base method.
func (m *MockCustomerUsecaseInterface) UpdateByID(customer *customermodel.Customer) (*customermodel.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", customer)
	ret0, _ := ret[0].(*customermodel.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockCustomerUsecaseInterfaceMockRecorder) UpdateByID(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockCustomerUsecaseInterface)(nil).UpdateByID), customer)
}