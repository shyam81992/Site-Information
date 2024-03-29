// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shyam81992/Site-Information/camqp (interfaces: ICAMQP,ICAMQPConn,ICAMQPChannel)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	camqp "github.com/shyam81992/Site-Information/camqp"
	amqp "github.com/streadway/amqp"
)

// MockICAMQP is a mock of ICAMQP interface.
type MockICAMQP struct {
	ctrl     *gomock.Controller
	recorder *MockICAMQPMockRecorder
}

// MockICAMQPMockRecorder is the mock recorder for MockICAMQP.
type MockICAMQPMockRecorder struct {
	mock *MockICAMQP
}

// NewMockICAMQP creates a new mock instance.
func NewMockICAMQP(ctrl *gomock.Controller) *MockICAMQP {
	mock := &MockICAMQP{ctrl: ctrl}
	mock.recorder = &MockICAMQPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICAMQP) EXPECT() *MockICAMQPMockRecorder {
	return m.recorder
}

// Publishmsg mocks base method.
func (m *MockICAMQP) Publishmsg(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publishmsg", arg0)
}

// Publishmsg indicates an expected call of Publishmsg.
func (mr *MockICAMQPMockRecorder) Publishmsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publishmsg", reflect.TypeOf((*MockICAMQP)(nil).Publishmsg), arg0)
}

// MockICAMQPConn is a mock of ICAMQPConn interface.
type MockICAMQPConn struct {
	ctrl     *gomock.Controller
	recorder *MockICAMQPConnMockRecorder
}

// MockICAMQPConnMockRecorder is the mock recorder for MockICAMQPConn.
type MockICAMQPConnMockRecorder struct {
	mock *MockICAMQPConn
}

// NewMockICAMQPConn creates a new mock instance.
func NewMockICAMQPConn(ctrl *gomock.Controller) *MockICAMQPConn {
	mock := &MockICAMQPConn{ctrl: ctrl}
	mock.recorder = &MockICAMQPConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICAMQPConn) EXPECT() *MockICAMQPConnMockRecorder {
	return m.recorder
}

// Channel mocks base method.
func (m *MockICAMQPConn) Channel() (camqp.ICAMQPChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel")
	ret0, _ := ret[0].(camqp.ICAMQPChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel.
func (mr *MockICAMQPConnMockRecorder) Channel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockICAMQPConn)(nil).Channel))
}

// Close mocks base method.
func (m *MockICAMQPConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockICAMQPConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockICAMQPConn)(nil).Close))
}

// MockICAMQPChannel is a mock of ICAMQPChannel interface.
type MockICAMQPChannel struct {
	ctrl     *gomock.Controller
	recorder *MockICAMQPChannelMockRecorder
}

// MockICAMQPChannelMockRecorder is the mock recorder for MockICAMQPChannel.
type MockICAMQPChannelMockRecorder struct {
	mock *MockICAMQPChannel
}

// NewMockICAMQPChannel creates a new mock instance.
func NewMockICAMQPChannel(ctrl *gomock.Controller) *MockICAMQPChannel {
	mock := &MockICAMQPChannel{ctrl: ctrl}
	mock.recorder = &MockICAMQPChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICAMQPChannel) EXPECT() *MockICAMQPChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockICAMQPChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockICAMQPChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockICAMQPChannel)(nil).Close))
}

// Publish mocks base method.
func (m *MockICAMQPChannel) Publish(arg0, arg1 string, arg2, arg3 bool, arg4 amqp.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockICAMQPChannelMockRecorder) Publish(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockICAMQPChannel)(nil).Publish), arg0, arg1, arg2, arg3, arg4)
}
