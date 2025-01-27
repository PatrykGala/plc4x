/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.28.1. DO NOT EDIT.

package spi

import mock "github.com/stretchr/testify/mock"

// MockRequestTransactionManager is an autogenerated mock type for the RequestTransactionManager type
type MockRequestTransactionManager struct {
	mock.Mock
}

type MockRequestTransactionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestTransactionManager) EXPECT() *MockRequestTransactionManager_Expecter {
	return &MockRequestTransactionManager_Expecter{mock: &_m.Mock}
}

// SetNumberOfConcurrentRequests provides a mock function with given fields: numberOfConcurrentRequests
func (_m *MockRequestTransactionManager) SetNumberOfConcurrentRequests(numberOfConcurrentRequests int) {
	_m.Called(numberOfConcurrentRequests)
}

// MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNumberOfConcurrentRequests'
type MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call struct {
	*mock.Call
}

// SetNumberOfConcurrentRequests is a helper method to define mock.On call
//   - numberOfConcurrentRequests int
func (_e *MockRequestTransactionManager_Expecter) SetNumberOfConcurrentRequests(numberOfConcurrentRequests interface{}) *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call {
	return &MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call{Call: _e.mock.On("SetNumberOfConcurrentRequests", numberOfConcurrentRequests)}
}

func (_c *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call) Run(run func(numberOfConcurrentRequests int)) *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call) Return() *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call) RunAndReturn(run func(int)) *MockRequestTransactionManager_SetNumberOfConcurrentRequests_Call {
	_c.Call.Return(run)
	return _c
}

// StartTransaction provides a mock function with given fields:
func (_m *MockRequestTransactionManager) StartTransaction() RequestTransaction {
	ret := _m.Called()

	var r0 RequestTransaction
	if rf, ok := ret.Get(0).(func() RequestTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RequestTransaction)
		}
	}

	return r0
}

// MockRequestTransactionManager_StartTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartTransaction'
type MockRequestTransactionManager_StartTransaction_Call struct {
	*mock.Call
}

// StartTransaction is a helper method to define mock.On call
func (_e *MockRequestTransactionManager_Expecter) StartTransaction() *MockRequestTransactionManager_StartTransaction_Call {
	return &MockRequestTransactionManager_StartTransaction_Call{Call: _e.mock.On("StartTransaction")}
}

func (_c *MockRequestTransactionManager_StartTransaction_Call) Run(run func()) *MockRequestTransactionManager_StartTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequestTransactionManager_StartTransaction_Call) Return(_a0 RequestTransaction) *MockRequestTransactionManager_StartTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransactionManager_StartTransaction_Call) RunAndReturn(run func() RequestTransaction) *MockRequestTransactionManager_StartTransaction_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRequestTransactionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRequestTransactionManager creates a new instance of MockRequestTransactionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRequestTransactionManager(t mockConstructorTestingTNewMockRequestTransactionManager) *MockRequestTransactionManager {
	mock := &MockRequestTransactionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
