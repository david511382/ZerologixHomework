// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/bll/trade/interface.go

// Package trade is a generated GoMock package.
package trade

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITradeLogic is a mock of ITradeLogic interface.
type MockITradeLogic struct {
	ctrl     *gomock.Controller
	recorder *MockITradeLogicMockRecorder
}

// MockITradeLogicMockRecorder is the mock recorder for MockITradeLogic.
type MockITradeLogicMockRecorder struct {
	mock *MockITradeLogic
}

// NewMockITradeLogic creates a new mock instance.
func NewMockITradeLogic(ctrl *gomock.Controller) *MockITradeLogic {
	mock := &MockITradeLogic{ctrl: ctrl}
	mock.recorder = &MockITradeLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITradeLogic) EXPECT() *MockITradeLogicMockRecorder {
	return m.recorder
}

// Match mocks base method.
func (m *MockITradeLogic) Match(order Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockITradeLogicMockRecorder) Match(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockITradeLogic)(nil).Match), order)
}

// MockITradeStorage is a mock of ITradeStorage interface.
type MockITradeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockITradeStorageMockRecorder
}

// MockITradeStorageMockRecorder is the mock recorder for MockITradeStorage.
type MockITradeStorageMockRecorder struct {
	mock *MockITradeStorage
}

// NewMockITradeStorage creates a new mock instance.
func NewMockITradeStorage(ctrl *gomock.Controller) *MockITradeStorage {
	mock := &MockITradeStorage{ctrl: ctrl}
	mock.recorder = &MockITradeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITradeStorage) EXPECT() *MockITradeStorageMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockITradeStorage) Delete(price float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", price)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITradeStorageMockRecorder) Delete(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITradeStorage)(nil).Delete), price)
}

// Load mocks base method.
func (m *MockITradeStorage) Load(price float64) ([]*Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", price)
	ret0, _ := ret[0].([]*Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockITradeStorageMockRecorder) Load(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockITradeStorage)(nil).Load), price)
}

// LoadID mocks base method.
func (m *MockITradeStorage) LoadID() (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadID")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadID indicates an expected call of LoadID.
func (mr *MockITradeStorageMockRecorder) LoadID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadID", reflect.TypeOf((*MockITradeStorage)(nil).LoadID))
}

// LockID mocks base method.
func (m *MockITradeStorage) LockID() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockID")
	ret0, _ := ret[0].(error)
	return ret0
}

// LockID indicates an expected call of LockID.
func (mr *MockITradeStorageMockRecorder) LockID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockID", reflect.TypeOf((*MockITradeStorage)(nil).LockID))
}

// LockOrders mocks base method.
func (m *MockITradeStorage) LockOrders(price float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockOrders", price)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockOrders indicates an expected call of LockOrders.
func (mr *MockITradeStorageMockRecorder) LockOrders(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockOrders", reflect.TypeOf((*MockITradeStorage)(nil).LockOrders), price)
}

// ReleaseIDLock mocks base method.
func (m *MockITradeStorage) ReleaseIDLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseIDLock")
}

// ReleaseIDLock indicates an expected call of ReleaseIDLock.
func (mr *MockITradeStorageMockRecorder) ReleaseIDLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseIDLock", reflect.TypeOf((*MockITradeStorage)(nil).ReleaseIDLock))
}

// ReleaseOrdersLock mocks base method.
func (m *MockITradeStorage) ReleaseOrdersLock(price float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseOrdersLock", price)
}

// ReleaseOrdersLock indicates an expected call of ReleaseOrdersLock.
func (mr *MockITradeStorageMockRecorder) ReleaseOrdersLock(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseOrdersLock", reflect.TypeOf((*MockITradeStorage)(nil).ReleaseOrdersLock), price)
}

// Set mocks base method.
func (m *MockITradeStorage) Set(price float64, orders []*Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", price, orders)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockITradeStorageMockRecorder) Set(price, orders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockITradeStorage)(nil).Set), price, orders)
}

// UpdateID mocks base method.
func (m *MockITradeStorage) UpdateID(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateID indicates an expected call of UpdateID.
func (mr *MockITradeStorageMockRecorder) UpdateID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateID", reflect.TypeOf((*MockITradeStorage)(nil).UpdateID), id)
}
