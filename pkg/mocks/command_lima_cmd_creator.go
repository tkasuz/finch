// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch/pkg/command (interfaces: LimaCmdCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	command "github.com/runfinch/finch/pkg/command"
)

// LimaCmdCreator is a mock of LimaCmdCreator interface.
type LimaCmdCreator struct {
	ctrl     *gomock.Controller
	recorder *LimaCmdCreatorMockRecorder
}

// LimaCmdCreatorMockRecorder is the mock recorder for LimaCmdCreator.
type LimaCmdCreatorMockRecorder struct {
	mock *LimaCmdCreator
}

// NewLimaCmdCreator creates a new mock instance.
func NewLimaCmdCreator(ctrl *gomock.Controller) *LimaCmdCreator {
	mock := &LimaCmdCreator{ctrl: ctrl}
	mock.recorder = &LimaCmdCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *LimaCmdCreator) EXPECT() *LimaCmdCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *LimaCmdCreator) Create(arg0 ...string) command.Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(command.Command)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *LimaCmdCreatorMockRecorder) Create(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*LimaCmdCreator)(nil).Create), arg0...)
}

// CreateWithoutStdio mocks base method.
func (m *LimaCmdCreator) CreateWithoutStdio(arg0 ...string) command.Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWithoutStdio", varargs...)
	ret0, _ := ret[0].(command.Command)
	return ret0
}

// CreateWithoutStdio indicates an expected call of CreateWithoutStdio.
func (mr *LimaCmdCreatorMockRecorder) CreateWithoutStdio(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithoutStdio", reflect.TypeOf((*LimaCmdCreator)(nil).CreateWithoutStdio), arg0...)
}

// RunWithReplacingStdout mocks base method.
func (m *LimaCmdCreator) RunWithReplacingStdout(arg0 []command.Replacement, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunWithReplacingStdout", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWithReplacingStdout indicates an expected call of RunWithReplacingStdout.
func (mr *LimaCmdCreatorMockRecorder) RunWithReplacingStdout(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWithReplacingStdout", reflect.TypeOf((*LimaCmdCreator)(nil).RunWithReplacingStdout), varargs...)
}
