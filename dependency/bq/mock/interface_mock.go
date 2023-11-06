// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package interface_mock is a generated GoMock package.
package interface_mock

import (
	context "context"
	reflect "reflect"

	bq "github.com/gasprawira/tokopedia-library/dependency/bq"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// IsClientNil mocks base method.
func (m *MockClient) IsClientNil() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClientNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClientNil indicates an expected call of IsClientNil.
func (mr *MockClientMockRecorder) IsClientNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClientNil", reflect.TypeOf((*MockClient)(nil).IsClientNil))
}

// Query mocks base method.
func (m *MockClient) Query(q string) bq.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", q)
	ret0, _ := ret[0].(bq.Query)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockClientMockRecorder) Query(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClient)(nil).Query), q)
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

// Read mocks base method.
func (m *MockQuery) Read(ctx context.Context) (bq.RowIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(bq.RowIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockQueryMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockQuery)(nil).Read), ctx)
}

// MockRowIterator is a mock of RowIterator interface.
type MockRowIterator struct {
	ctrl     *gomock.Controller
	recorder *MockRowIteratorMockRecorder
}

// MockRowIteratorMockRecorder is the mock recorder for MockRowIterator.
type MockRowIteratorMockRecorder struct {
	mock *MockRowIterator
}

// NewMockRowIterator creates a new mock instance.
func NewMockRowIterator(ctrl *gomock.Controller) *MockRowIterator {
	mock := &MockRowIterator{ctrl: ctrl}
	mock.recorder = &MockRowIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRowIterator) EXPECT() *MockRowIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockRowIterator) Next(dst interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowIteratorMockRecorder) Next(dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRowIterator)(nil).Next), dst)
}