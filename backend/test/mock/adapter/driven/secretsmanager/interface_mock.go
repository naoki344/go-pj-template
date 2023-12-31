// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/naoki.miyoshi/Data/pjct/en/en/backend/test/scripts/../../internal/adapter/driven/secretsmanager/interface.go

// Package mock_secretsmanager is a generated GoMock package.
package mock_secretsmanager

import (
	context "context"
	reflect "reflect"

	secretsmanager "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsManagerClientInterface is a mock of SecretsManagerClientInterface interface.
type MockSecretsManagerClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerClientInterfaceMockRecorder
}

// MockSecretsManagerClientInterfaceMockRecorder is the mock recorder for MockSecretsManagerClientInterface.
type MockSecretsManagerClientInterfaceMockRecorder struct {
	mock *MockSecretsManagerClientInterface
}

// NewMockSecretsManagerClientInterface creates a new mock instance.
func NewMockSecretsManagerClientInterface(ctrl *gomock.Controller) *MockSecretsManagerClientInterface {
	mock := &MockSecretsManagerClientInterface{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerClientInterface) EXPECT() *MockSecretsManagerClientInterfaceMockRecorder {
	return m.recorder
}

// GetSecretStringWithContext mocks base method.
func (m *MockSecretsManagerClientInterface) GetSecretStringWithContext(ctx context.Context, secretsID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretStringWithContext", ctx, secretsID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretStringWithContext indicates an expected call of GetSecretStringWithContext.
func (mr *MockSecretsManagerClientInterfaceMockRecorder) GetSecretStringWithContext(ctx, secretsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretStringWithContext", reflect.TypeOf((*MockSecretsManagerClientInterface)(nil).GetSecretStringWithContext), ctx, secretsID)
}

// MockSecretsManagerAdapterInterface is a mock of SecretsManagerAdapterInterface interface.
type MockSecretsManagerAdapterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerAdapterInterfaceMockRecorder
}

// MockSecretsManagerAdapterInterfaceMockRecorder is the mock recorder for MockSecretsManagerAdapterInterface.
type MockSecretsManagerAdapterInterfaceMockRecorder struct {
	mock *MockSecretsManagerAdapterInterface
}

// NewMockSecretsManagerAdapterInterface creates a new mock instance.
func NewMockSecretsManagerAdapterInterface(ctrl *gomock.Controller) *MockSecretsManagerAdapterInterface {
	mock := &MockSecretsManagerAdapterInterface{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerAdapterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerAdapterInterface) EXPECT() *MockSecretsManagerAdapterInterfaceMockRecorder {
	return m.recorder
}

// GetPrimaryDBAccount mocks base method.
func (m *MockSecretsManagerAdapterInterface) GetPrimaryDBAccount(ctx context.Context) secretsmanager.DBAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrimaryDBAccount", ctx)
	ret0, _ := ret[0].(secretsmanager.DBAccount)
	return ret0
}

// GetPrimaryDBAccount indicates an expected call of GetPrimaryDBAccount.
func (mr *MockSecretsManagerAdapterInterfaceMockRecorder) GetPrimaryDBAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrimaryDBAccount", reflect.TypeOf((*MockSecretsManagerAdapterInterface)(nil).GetPrimaryDBAccount), ctx)
}

// GetSecondaryDBAccount mocks base method.
func (m *MockSecretsManagerAdapterInterface) GetSecondaryDBAccount(ctx context.Context) secretsmanager.DBAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecondaryDBAccount", ctx)
	ret0, _ := ret[0].(secretsmanager.DBAccount)
	return ret0
}

// GetSecondaryDBAccount indicates an expected call of GetSecondaryDBAccount.
func (mr *MockSecretsManagerAdapterInterfaceMockRecorder) GetSecondaryDBAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecondaryDBAccount", reflect.TypeOf((*MockSecretsManagerAdapterInterface)(nil).GetSecondaryDBAccount), ctx)
}
