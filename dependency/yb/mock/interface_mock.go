// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package interface_mock is a generated GoMock package.
package interface_mock

import (
	reflect "reflect"

	yb "github.com/gasprawira/tokopedia-library/dependency/yb"
	gomock "github.com/golang/mock/gomock"
)

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Batch mocks base method.
func (m *MockSession) Batch(batchType yb.BatchType) yb.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Batch", batchType)
	ret0, _ := ret[0].(yb.Batch)
	return ret0
}

// Batch indicates an expected call of Batch.
func (mr *MockSessionMockRecorder) Batch(batchType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockSession)(nil).Batch), batchType)
}

// IsClientNil mocks base method.
func (m *MockSession) IsClientNil() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClientNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClientNil indicates an expected call of IsClientNil.
func (mr *MockSessionMockRecorder) IsClientNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClientNil", reflect.TypeOf((*MockSession)(nil).IsClientNil))
}

// Query mocks base method.
func (m *MockSession) Query(statement string, arguments []interface{}) yb.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", statement, arguments)
	ret0, _ := ret[0].(yb.Query)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockSessionMockRecorder) Query(statement, arguments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSession)(nil).Query), statement, arguments)
}

// MockQuery is a mock of Query interface.
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance.
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockQuery) Exec() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec")
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockQueryMockRecorder) Exec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockQuery)(nil).Exec))
}

// Iter mocks base method.
func (m *MockQuery) Iter() yb.Iter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(yb.Iter)
	return ret0
}

// Iter indicates an expected call of Iter.
func (mr *MockQueryMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockQuery)(nil).Iter))
}

// MockIter is a mock of Iter interface.
type MockIter struct {
	ctrl     *gomock.Controller
	recorder *MockIterMockRecorder
}

// MockIterMockRecorder is the mock recorder for MockIter.
type MockIterMockRecorder struct {
	mock *MockIter
}

// NewMockIter creates a new mock instance.
func NewMockIter(ctrl *gomock.Controller) *MockIter {
	mock := &MockIter{ctrl: ctrl}
	mock.recorder = &MockIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIter) EXPECT() *MockIterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIter)(nil).Close))
}

// MapScan mocks base method.
func (m *MockIter) MapScan(arg0 map[string]interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapScan", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MapScan indicates an expected call of MapScan.
func (mr *MockIterMockRecorder) MapScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapScan", reflect.TypeOf((*MockIter)(nil).MapScan), arg0)
}

// MockBatch is a mock of Batch interface.
type MockBatch struct {
	ctrl     *gomock.Controller
	recorder *MockBatchMockRecorder
}

// MockBatchMockRecorder is the mock recorder for MockBatch.
type MockBatchMockRecorder struct {
	mock *MockBatch
}

// NewMockBatch creates a new mock instance.
func NewMockBatch(ctrl *gomock.Controller) *MockBatch {
	mock := &MockBatch{ctrl: ctrl}
	mock.recorder = &MockBatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatch) EXPECT() *MockBatchMockRecorder {
	return m.recorder
}

// AddQuery mocks base method.
func (m *MockBatch) AddQuery(statement string, arguments ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{statement}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddQuery", varargs...)
}

// AddQuery indicates an expected call of AddQuery.
func (mr *MockBatchMockRecorder) AddQuery(statement interface{}, arguments ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{statement}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQuery", reflect.TypeOf((*MockBatch)(nil).AddQuery), varargs...)
}

// ExecuteBatch mocks base method.
func (m *MockBatch) ExecuteBatch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteBatch")
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteBatch indicates an expected call of ExecuteBatch.
func (mr *MockBatchMockRecorder) ExecuteBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteBatch", reflect.TypeOf((*MockBatch)(nil).ExecuteBatch))
}