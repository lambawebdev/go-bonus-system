// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lambawebdev/go-bonus-system/internal/gophemart/repositories (interfaces: OrderRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

// MockOrderRepo is a mock of OrderRepo interface.
type MockOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoMockRecorder
}

// MockOrderRepoMockRecorder is the mock recorder for MockOrderRepo.
type MockOrderRepoMockRecorder struct {
	mock *MockOrderRepo
}

// NewMockOrderRepo creates a new mock instance.
func NewMockOrderRepo(ctrl *gomock.Controller) *MockOrderRepo {
	mock := &MockOrderRepo{ctrl: ctrl}
	mock.recorder = &MockOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepo) EXPECT() *MockOrderRepoMockRecorder {
	return m.recorder
}

// CheckIfOrderWasAddedByAnotherUser mocks base method.
func (m *MockOrderRepo) CheckIfOrderWasAddedByAnotherUser(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfOrderWasAddedByAnotherUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfOrderWasAddedByAnotherUser indicates an expected call of CheckIfOrderWasAddedByAnotherUser.
func (mr *MockOrderRepoMockRecorder) CheckIfOrderWasAddedByAnotherUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfOrderWasAddedByAnotherUser", reflect.TypeOf((*MockOrderRepo)(nil).CheckIfOrderWasAddedByAnotherUser), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockOrderRepo) CreateOrder(arg0 context.Context, arg1 string) (entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepoMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepo)(nil).CreateOrder), arg0, arg1)
}

// GetNotProcessedOrders mocks base method.
func (m *MockOrderRepo) GetNotProcessedOrders() ([]entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotProcessedOrders")
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotProcessedOrders indicates an expected call of GetNotProcessedOrders.
func (mr *MockOrderRepoMockRecorder) GetNotProcessedOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotProcessedOrders", reflect.TypeOf((*MockOrderRepo)(nil).GetNotProcessedOrders))
}

// GetOrderByNumber mocks base method.
func (m *MockOrderRepo) GetOrderByNumber(arg0 string) (entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByNumber", arg0)
	ret0, _ := ret[0].(entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByNumber indicates an expected call of GetOrderByNumber.
func (mr *MockOrderRepoMockRecorder) GetOrderByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByNumber", reflect.TypeOf((*MockOrderRepo)(nil).GetOrderByNumber), arg0)
}

// GetOrders mocks base method.
func (m *MockOrderRepo) GetOrders(arg0 context.Context) ([]entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepoMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepo)(nil).GetOrders), arg0)
}

// UpdateOrderStatus mocks base method.
func (m *MockOrderRepo) UpdateOrderStatus(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockOrderRepoMockRecorder) UpdateOrderStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockOrderRepo)(nil).UpdateOrderStatus), arg0, arg1)
}