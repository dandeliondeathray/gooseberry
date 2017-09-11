// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dandeliondeathray/gooseberry/director (interfaces: Work)

package director_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWork is a mock of Work interface
type MockWork struct {
	ctrl     *gomock.Controller
	recorder *MockWorkMockRecorder
}

// MockWorkMockRecorder is the mock recorder for MockWork
type MockWorkMockRecorder struct {
	mock *MockWork
}

// NewMockWork creates a new mock instance
func NewMockWork(ctrl *gomock.Controller) *MockWork {
	mock := &MockWork{ctrl: ctrl}
	mock.recorder = &MockWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWork) EXPECT() *MockWorkMockRecorder {
	return m.recorder
}

// Schedule mocks base method
func (m *MockWork) Schedule() {
	m.ctrl.Call(m, "Schedule")
}

// Schedule indicates an expected call of Schedule
func (mr *MockWorkMockRecorder) Schedule() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockWork)(nil).Schedule))
}
