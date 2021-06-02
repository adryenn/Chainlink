// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BlockFetcherConfig is an autogenerated mock type for the BlockFetcherConfig type
type BlockFetcherConfig struct {
	mock.Mock
}

// BlockBackfillDepth provides a mock function with given fields:
func (_m *BlockFetcherConfig) BlockBackfillDepth() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// BlockFetcherBatchSize provides a mock function with given fields:
func (_m *BlockFetcherConfig) BlockFetcherBatchSize() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// EthFinalityDepth provides a mock function with given fields:
func (_m *BlockFetcherConfig) EthFinalityDepth() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EthHeadTrackerHistoryDepth provides a mock function with given fields:
func (_m *BlockFetcherConfig) EthHeadTrackerHistoryDepth() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
