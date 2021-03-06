//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/uber/aresdb/memstore/common"

import mock "github.com/stretchr/testify/mock"

// sortedColumnIterator is an autogenerated mock type for the sortedColumnIterator type
type sortedColumnIterator struct {
	mock.Mock
}

// count provides a mock function with given fields:
func (_m *sortedColumnIterator) count() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// currentPosition provides a mock function with given fields:
func (_m *sortedColumnIterator) currentPosition() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// currentSkipRows provides a mock function with given fields:
func (_m *sortedColumnIterator) currentSkipRows() []int {
	ret := _m.Called()

	var r0 []int
	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// done provides a mock function with given fields:
func (_m *sortedColumnIterator) done() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// index provides a mock function with given fields:
func (_m *sortedColumnIterator) index() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// next provides a mock function with given fields:
func (_m *sortedColumnIterator) next() {
	_m.Called()
}

// nextPosition provides a mock function with given fields:
func (_m *sortedColumnIterator) nextPosition() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// read provides a mock function with given fields:
func (_m *sortedColumnIterator) read() {
	_m.Called()
}

// setEndPosition provides a mock function with given fields: pos
func (_m *sortedColumnIterator) setEndPosition(pos uint32) {
	_m.Called(pos)
}

// value provides a mock function with given fields:
func (_m *sortedColumnIterator) value() common.DataValue {
	ret := _m.Called()

	var r0 common.DataValue
	if rf, ok := ret.Get(0).(func() common.DataValue); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.DataValue)
	}

	return r0
}
