// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch/pkg/fssh (interfaces: Dialer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ssh "golang.org/x/crypto/ssh"
)

// Dialer is a mock of Dialer interface.
type Dialer struct {
	ctrl     *gomock.Controller
	recorder *DialerMockRecorder
}

// DialerMockRecorder is the mock recorder for Dialer.
type DialerMockRecorder struct {
	mock *Dialer
}

// NewDialer creates a new mock instance.
func NewDialer(ctrl *gomock.Controller) *Dialer {
	mock := &Dialer{ctrl: ctrl}
	mock.recorder = &DialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Dialer) EXPECT() *DialerMockRecorder {
	return m.recorder
}

// Dial mocks base method.
func (m *Dialer) Dial(arg0, arg1 string, arg2 *ssh.ClientConfig) (*ssh.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dial", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ssh.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dial indicates an expected call of Dial.
func (mr *DialerMockRecorder) Dial(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*Dialer)(nil).Dial), arg0, arg1, arg2)
}
