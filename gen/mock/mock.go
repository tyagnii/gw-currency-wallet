// Code generated by MockGen. DO NOT EDIT.
// Source: internal/db/db.go
//
// Generated by this command:
//
//	mockgen -source internal/db/db.go -destination gen/mock/mock.go
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	models "github.com/tyagnii/gw-currency-wallet/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockDBConnector is a mock of DBConnector interface.
type MockDBConnector struct {
	ctrl     *gomock.Controller
	recorder *MockDBConnectorMockRecorder
}

// MockDBConnectorMockRecorder is the mock recorder for MockDBConnector.
type MockDBConnectorMockRecorder struct {
	mock *MockDBConnector
}

// NewMockDBConnector creates a new mock instance.
func NewMockDBConnector(ctrl *gomock.Controller) *MockDBConnector {
	mock := &MockDBConnector{ctrl: ctrl}
	mock.recorder = &MockDBConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBConnector) EXPECT() *MockDBConnectorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockDBConnector) CreateUser(ctx context.Context, u models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDBConnectorMockRecorder) CreateUser(ctx, u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDBConnector)(nil).CreateUser), ctx, u)
}

// CreateWallet mocks base method.
func (m *MockDBConnector) CreateWallet(ctx context.Context, w models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", ctx, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockDBConnectorMockRecorder) CreateWallet(ctx, w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockDBConnector)(nil).CreateWallet), ctx, w)
}

// Deposit mocks base method.
func (m *MockDBConnector) Deposit(ctx context.Context, w models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", ctx, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deposit indicates an expected call of Deposit.
func (mr *MockDBConnectorMockRecorder) Deposit(ctx, w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockDBConnector)(nil).Deposit), ctx, w)
}

// GetBalance mocks base method.
func (m *MockDBConnector) GetBalance(ctx context.Context, u models.User) (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, u)
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockDBConnectorMockRecorder) GetBalance(ctx, u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockDBConnector)(nil).GetBalance), ctx, u)
}

// GetUser mocks base method.
func (m *MockDBConnector) GetUser(ctx context.Context, u models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, u)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDBConnectorMockRecorder) GetUser(ctx, u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDBConnector)(nil).GetUser), ctx, u)
}

// Withdraw mocks base method.
func (m *MockDBConnector) Withdraw(ctx context.Context, w models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockDBConnectorMockRecorder) Withdraw(ctx, w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockDBConnector)(nil).Withdraw), ctx, w)
}