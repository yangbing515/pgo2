// Code generated by MockGen. DO NOT EDIT.
// Source: config/interface.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/pinguo/pgo2/config"
	reflect "reflect"
)

// MockIConfigParser is a mock of IConfigParser interface
type MockIConfigParser struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigParserMockRecorder
}

// MockIConfigParserMockRecorder is the mock recorder for MockIConfigParser
type MockIConfigParserMockRecorder struct {
	mock *MockIConfigParser
}

// NewMockIConfigParser creates a new mock instance
func NewMockIConfigParser(ctrl *gomock.Controller) *MockIConfigParser {
	mock := &MockIConfigParser{ctrl: ctrl}
	mock.recorder = &MockIConfigParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConfigParser) EXPECT() *MockIConfigParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockIConfigParser) Parse(path string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", path)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockIConfigParserMockRecorder) Parse(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockIConfigParser)(nil).Parse), path)
}

// MockIConfig is a mock of IConfig interface
type MockIConfig struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigMockRecorder
}

// MockIConfigMockRecorder is the mock recorder for MockIConfig
type MockIConfigMockRecorder struct {
	mock *MockIConfig
}

// NewMockIConfig creates a new mock instance
func NewMockIConfig(ctrl *gomock.Controller) *MockIConfig {
	mock := &MockIConfig{ctrl: ctrl}
	mock.recorder = &MockIConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConfig) EXPECT() *MockIConfigMockRecorder {
	return m.recorder
}

// AddParser mocks base method
func (m *MockIConfig) AddParser(ext string, parser config.IConfigParser) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddParser", ext, parser)
}

// AddParser indicates an expected call of AddParser
func (mr *MockIConfigMockRecorder) AddParser(ext, parser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddParser", reflect.TypeOf((*MockIConfig)(nil).AddParser), ext, parser)
}

// AddPath mocks base method
func (m *MockIConfig) AddPath(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPath", path)
}

// AddPath indicates an expected call of AddPath
func (mr *MockIConfigMockRecorder) AddPath(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPath", reflect.TypeOf((*MockIConfig)(nil).AddPath), path)
}

// GetBool mocks base method
func (m *MockIConfig) GetBool(key string, dft bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBool", key, dft)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool
func (mr *MockIConfigMockRecorder) GetBool(key, dft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockIConfig)(nil).GetBool), key, dft)
}

// GetInt mocks base method
func (m *MockIConfig) GetInt(key string, dft int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", key, dft)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt
func (mr *MockIConfigMockRecorder) GetInt(key, dft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockIConfig)(nil).GetInt), key, dft)
}

// GetFloat mocks base method
func (m *MockIConfig) GetFloat(key string, dft float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloat", key, dft)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetFloat indicates an expected call of GetFloat
func (mr *MockIConfigMockRecorder) GetFloat(key, dft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat", reflect.TypeOf((*MockIConfig)(nil).GetFloat), key, dft)
}

// GetString mocks base method
func (m *MockIConfig) GetString(key, dft string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", key, dft)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString
func (mr *MockIConfigMockRecorder) GetString(key, dft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockIConfig)(nil).GetString), key, dft)
}

// GetSliceBool mocks base method
func (m *MockIConfig) GetSliceBool(key string) []bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliceBool", key)
	ret0, _ := ret[0].([]bool)
	return ret0
}

// GetSliceBool indicates an expected call of GetSliceBool
func (mr *MockIConfigMockRecorder) GetSliceBool(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliceBool", reflect.TypeOf((*MockIConfig)(nil).GetSliceBool), key)
}

// GetSliceInt mocks base method
func (m *MockIConfig) GetSliceInt(key string) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliceInt", key)
	ret0, _ := ret[0].([]int)
	return ret0
}

// GetSliceInt indicates an expected call of GetSliceInt
func (mr *MockIConfigMockRecorder) GetSliceInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliceInt", reflect.TypeOf((*MockIConfig)(nil).GetSliceInt), key)
}

// GetSliceFloat mocks base method
func (m *MockIConfig) GetSliceFloat(key string) []float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliceFloat", key)
	ret0, _ := ret[0].([]float64)
	return ret0
}

// GetSliceFloat indicates an expected call of GetSliceFloat
func (mr *MockIConfigMockRecorder) GetSliceFloat(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliceFloat", reflect.TypeOf((*MockIConfig)(nil).GetSliceFloat), key)
}

// GetSliceString mocks base method
func (m *MockIConfig) GetSliceString(key string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliceString", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSliceString indicates an expected call of GetSliceString
func (mr *MockIConfigMockRecorder) GetSliceString(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliceString", reflect.TypeOf((*MockIConfig)(nil).GetSliceString), key)
}

// Get mocks base method
func (m *MockIConfig) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockIConfigMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIConfig)(nil).Get), key)
}

// Set mocks base method
func (m *MockIConfig) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set
func (mr *MockIConfigMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIConfig)(nil).Set), key, val)
}

// CheckPath mocks base method
func (m *MockIConfig) CheckPath() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckPath")
}

// CheckPath indicates an expected call of CheckPath
func (mr *MockIConfigMockRecorder) CheckPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPath", reflect.TypeOf((*MockIConfig)(nil).CheckPath))
}
