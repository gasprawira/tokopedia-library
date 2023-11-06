// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package interface_mock is a generated GoMock package.
package interface_mock

import (
	reflect "reflect"

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

// Command mocks base method.
func (m *MockClient) Command(command string, args ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{command}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Command", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Command indicates an expected call of Command.
func (mr *MockClientMockRecorder) Command(command interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{command}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockClient)(nil).Command), varargs...)
}

// DecrBy mocks base method.
func (m *MockClient) DecrBy(key string, amount int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", key, amount)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockClientMockRecorder) DecrBy(key, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockClient)(nil).DecrBy), key, amount)
}

// Del mocks base method.
func (m *MockClient) Del(key string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del.
func (mr *MockClientMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockClient)(nil).Del), key)
}

// EXISTS mocks base method.
func (m *MockClient) EXISTS(key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EXISTS", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EXISTS indicates an expected call of EXISTS.
func (mr *MockClientMockRecorder) EXISTS(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EXISTS", reflect.TypeOf((*MockClient)(nil).EXISTS), key)
}

// Expire mocks base method.
func (m *MockClient) Expire(key string, ttl int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", key, ttl)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expire indicates an expected call of Expire.
func (mr *MockClientMockRecorder) Expire(key, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockClient)(nil).Expire), key, ttl)
}

// Get mocks base method.
func (m *MockClient) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), key)
}

// HExists mocks base method.
func (m *MockClient) HExists(key, field string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HExists", key, field)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HExists indicates an expected call of HExists.
func (mr *MockClientMockRecorder) HExists(key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HExists", reflect.TypeOf((*MockClient)(nil).HExists), key, field)
}

// HGet mocks base method.
func (m *MockClient) HGet(key, field string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", key, field)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGet indicates an expected call of HGet.
func (mr *MockClientMockRecorder) HGet(key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockClient)(nil).HGet), key, field)
}

// HGetAll mocks base method.
func (m *MockClient) HGetAll(key string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGetAll indicates an expected call of HGetAll.
func (mr *MockClientMockRecorder) HGetAll(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockClient)(nil).HGetAll), key)
}

// HMGet mocks base method.
func (m *MockClient) HMGet(key string, fields []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HMGet", key, fields)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HMGet indicates an expected call of HMGet.
func (mr *MockClientMockRecorder) HMGet(key, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockClient)(nil).HMGet), key, fields)
}

// HMSet mocks base method.
func (m_2 *MockClient) HMSet(key string, ttl int, m map[string]string) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "HMSet", key, ttl, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// HMSet indicates an expected call of HMSet.
func (mr *MockClientMockRecorder) HMSet(key, ttl, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockClient)(nil).HMSet), key, ttl, m)
}

// HMSetWithoutTTL mocks base method.
func (m_2 *MockClient) HMSetWithoutTTL(key string, m map[string]string) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "HMSetWithoutTTL", key, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// HMSetWithoutTTL indicates an expected call of HMSetWithoutTTL.
func (mr *MockClientMockRecorder) HMSetWithoutTTL(key, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSetWithoutTTL", reflect.TypeOf((*MockClient)(nil).HMSetWithoutTTL), key, m)
}

// HSet mocks base method.
func (m *MockClient) HSet(key, field, value string, ttl int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", key, field, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockClientMockRecorder) HSet(key, field, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockClient)(nil).HSet), key, field, value, ttl)
}

// IncrBy mocks base method.
func (m *MockClient) IncrBy(key string, amount int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", key, amount)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockClientMockRecorder) IncrBy(key, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockClient)(nil).IncrBy), key, amount)
}

// MGet mocks base method.
func (m *MockClient) MGet(key ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range key {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGet indicates an expected call of MGet.
func (mr *MockClientMockRecorder) MGet(key ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockClient)(nil).MGet), key...)
}

// Setex mocks base method.
func (m *MockClient) Setex(key, value string, ttl int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setex", key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setex indicates an expected call of Setex.
func (mr *MockClientMockRecorder) Setex(key, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setex", reflect.TypeOf((*MockClient)(nil).Setex), key, value, ttl)
}

// Setnx mocks base method.
func (m *MockClient) Setnx(key, value string, ttl int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setnx", key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setnx indicates an expected call of Setnx.
func (mr *MockClientMockRecorder) Setnx(key, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setnx", reflect.TypeOf((*MockClient)(nil).Setnx), key, value, ttl)
}

// ZAddPairs mocks base method.
func (m *MockClient) ZAddPairs(key string, pairs map[string]float64, ttl int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddPairs", key, pairs, ttl)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZAddPairs indicates an expected call of ZAddPairs.
func (mr *MockClientMockRecorder) ZAddPairs(key, pairs, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddPairs", reflect.TypeOf((*MockClient)(nil).ZAddPairs), key, pairs, ttl)
}

// ZRange mocks base method.
func (m *MockClient) ZRange(key string, start, stop int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRange", key, start, stop)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRange indicates an expected call of ZRange.
func (mr *MockClientMockRecorder) ZRange(key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRange", reflect.TypeOf((*MockClient)(nil).ZRange), key, start, stop)
}

// ZRevRange mocks base method.
func (m *MockClient) ZRevRange(key string, start, stop int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRange", key, start, stop)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRevRange indicates an expected call of ZRevRange.
func (mr *MockClientMockRecorder) ZRevRange(key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRange", reflect.TypeOf((*MockClient)(nil).ZRevRange), key, start, stop)
}