// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/uber/aresdb/memstore/common"
import io "io"
import mock "github.com/stretchr/testify/mock"
import "unsafe"

// LiveVectorParty is an autogenerated mock type for the LiveVectorParty type
type LiveVectorParty struct {
	mock.Mock
}

// Allocate provides a mock function with given fields: hasCount
func (_m *LiveVectorParty) Allocate(hasCount bool) {
	_m.Called(hasCount)
}

// Equals provides a mock function with given fields: other
func (_m *LiveVectorParty) Equals(other common.VectorParty) bool {
	ret := _m.Called(other)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.VectorParty) bool); ok {
		r0 = rf(other)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetBytes provides a mock function with given fields:
func (_m *LiveVectorParty) GetBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetDataType provides a mock function with given fields:
func (_m *LiveVectorParty) GetDataType() common.DataType {
	ret := _m.Called()

	var r0 common.DataType
	if rf, ok := ret.Get(0).(func() common.DataType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.DataType)
	}

	return r0
}

// GetDataValue provides a mock function with given fields: offset
func (_m *LiveVectorParty) GetDataValue(offset int) common.DataValue {
	ret := _m.Called(offset)

	var r0 common.DataValue
	if rf, ok := ret.Get(0).(func(int) common.DataValue); ok {
		r0 = rf(offset)
	} else {
		r0 = ret.Get(0).(common.DataValue)
	}

	return r0
}

// GetDataValueByRow provides a mock function with given fields: row
func (_m *LiveVectorParty) GetDataValueByRow(row int) common.DataValue {
	ret := _m.Called(row)

	var r0 common.DataValue
	if rf, ok := ret.Get(0).(func(int) common.DataValue); ok {
		r0 = rf(row)
	} else {
		r0 = ret.Get(0).(common.DataValue)
	}

	return r0
}

// GetLength provides a mock function with given fields:
func (_m *LiveVectorParty) GetLength() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetMinMaxValue provides a mock function with given fields:
func (_m *LiveVectorParty) GetMinMaxValue() (uint32, uint32) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func() uint32); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(uint32)
	}

	return r0, r1
}

// GetNonDefaultValueCount provides a mock function with given fields:
func (_m *LiveVectorParty) GetNonDefaultValueCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetValidity provides a mock function with given fields: offset
func (_m *LiveVectorParty) GetValidity(offset int) bool {
	ret := _m.Called(offset)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(offset)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetValue provides a mock function with given fields: offset
func (_m *LiveVectorParty) GetValue(offset int) (unsafe.Pointer, bool) {
	ret := _m.Called(offset)

	var r0 unsafe.Pointer
	if rf, ok := ret.Get(0).(func(int) unsafe.Pointer); ok {
		r0 = rf(offset)
	} else {
		r0 = ret.Get(0).(unsafe.Pointer)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(int) bool); ok {
		r1 = rf(offset)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Read provides a mock function with given fields: reader, serializer
func (_m *LiveVectorParty) Read(reader io.Reader, serializer common.VectorPartySerializer) error {
	ret := _m.Called(reader, serializer)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader, common.VectorPartySerializer) error); ok {
		r0 = rf(reader, serializer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SafeDestruct provides a mock function with given fields:
func (_m *LiveVectorParty) SafeDestruct() {
	_m.Called()
}

// SetBool provides a mock function with given fields: offset, val, valid
func (_m *LiveVectorParty) SetBool(offset int, val bool, valid bool) {
	_m.Called(offset, val, valid)
}

// SetDataValue provides a mock function with given fields: offset, value, countsUpdateMode, counts
func (_m *LiveVectorParty) SetDataValue(offset int, value common.DataValue, countsUpdateMode common.ValueCountsUpdateMode, counts ...uint32) {
	_va := make([]interface{}, len(counts))
	for _i := range counts {
		_va[_i] = counts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, offset, value, countsUpdateMode)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetGoValue provides a mock function with given fields: offset, val, valid
func (_m *LiveVectorParty) SetGoValue(offset int, val common.GoDataValue, valid bool) {
	_m.Called(offset, val, valid)
}

// SetValue provides a mock function with given fields: offset, val, valid
func (_m *LiveVectorParty) SetValue(offset int, val unsafe.Pointer, valid bool) {
	_m.Called(offset, val, valid)
}

// Slice provides a mock function with given fields: startRow, numRows
func (_m *LiveVectorParty) Slice(startRow int, numRows int) common.SlicedVector {
	ret := _m.Called(startRow, numRows)

	var r0 common.SlicedVector
	if rf, ok := ret.Get(0).(func(int, int) common.SlicedVector); ok {
		r0 = rf(startRow, numRows)
	} else {
		r0 = ret.Get(0).(common.SlicedVector)
	}

	return r0
}

// Write provides a mock function with given fields: writer
func (_m *LiveVectorParty) Write(writer io.Writer) error {
	ret := _m.Called(writer)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(writer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
