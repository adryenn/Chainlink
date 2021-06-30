// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	flags_wrapper "github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/flags_wrapper"

	generated "github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Flags is an autogenerated mock type for the FlagsInterface type
type Flags struct {
	mock.Mock
}

// AcceptOwnership provides a mock function with given fields: opts
func (_m *Flags) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddAccess provides a mock function with given fields: opts, _user
func (_m *Flags) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _user)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields:
func (_m *Flags) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// CheckEnabled provides a mock function with given fields: opts
func (_m *Flags) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	ret := _m.Called(opts)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) bool); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableAccessCheck provides a mock function with given fields: opts
func (_m *Flags) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableAccessCheck provides a mock function with given fields: opts
func (_m *Flags) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterAddedAccess provides a mock function with given fields: opts
func (_m *Flags) FilterAddedAccess(opts *bind.FilterOpts) (*flags_wrapper.FlagsAddedAccessIterator, error) {
	ret := _m.Called(opts)

	var r0 *flags_wrapper.FlagsAddedAccessIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *flags_wrapper.FlagsAddedAccessIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsAddedAccessIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterCheckAccessDisabled provides a mock function with given fields: opts
func (_m *Flags) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*flags_wrapper.FlagsCheckAccessDisabledIterator, error) {
	ret := _m.Called(opts)

	var r0 *flags_wrapper.FlagsCheckAccessDisabledIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *flags_wrapper.FlagsCheckAccessDisabledIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsCheckAccessDisabledIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterCheckAccessEnabled provides a mock function with given fields: opts
func (_m *Flags) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*flags_wrapper.FlagsCheckAccessEnabledIterator, error) {
	ret := _m.Called(opts)

	var r0 *flags_wrapper.FlagsCheckAccessEnabledIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *flags_wrapper.FlagsCheckAccessEnabledIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsCheckAccessEnabledIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterFlagLowered provides a mock function with given fields: opts, subject
func (_m *Flags) FilterFlagLowered(opts *bind.FilterOpts, subject []common.Address) (*flags_wrapper.FlagsFlagLoweredIterator, error) {
	ret := _m.Called(opts, subject)

	var r0 *flags_wrapper.FlagsFlagLoweredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *flags_wrapper.FlagsFlagLoweredIterator); ok {
		r0 = rf(opts, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsFlagLoweredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterFlagRaised provides a mock function with given fields: opts, subject
func (_m *Flags) FilterFlagRaised(opts *bind.FilterOpts, subject []common.Address) (*flags_wrapper.FlagsFlagRaisedIterator, error) {
	ret := _m.Called(opts, subject)

	var r0 *flags_wrapper.FlagsFlagRaisedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *flags_wrapper.FlagsFlagRaisedIterator); ok {
		r0 = rf(opts, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsFlagRaisedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferRequested provides a mock function with given fields: opts, from, to
func (_m *Flags) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*flags_wrapper.FlagsOwnershipTransferRequestedIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *flags_wrapper.FlagsOwnershipTransferRequestedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *flags_wrapper.FlagsOwnershipTransferRequestedIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsOwnershipTransferRequestedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, from, to
func (_m *Flags) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*flags_wrapper.FlagsOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *flags_wrapper.FlagsOwnershipTransferredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *flags_wrapper.FlagsOwnershipTransferredIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsOwnershipTransferredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRaisingAccessControllerUpdated provides a mock function with given fields: opts, previous, current
func (_m *Flags) FilterRaisingAccessControllerUpdated(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*flags_wrapper.FlagsRaisingAccessControllerUpdatedIterator, error) {
	ret := _m.Called(opts, previous, current)

	var r0 *flags_wrapper.FlagsRaisingAccessControllerUpdatedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *flags_wrapper.FlagsRaisingAccessControllerUpdatedIterator); ok {
		r0 = rf(opts, previous, current)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsRaisingAccessControllerUpdatedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, previous, current)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRemovedAccess provides a mock function with given fields: opts
func (_m *Flags) FilterRemovedAccess(opts *bind.FilterOpts) (*flags_wrapper.FlagsRemovedAccessIterator, error) {
	ret := _m.Called(opts)

	var r0 *flags_wrapper.FlagsRemovedAccessIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *flags_wrapper.FlagsRemovedAccessIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsRemovedAccessIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlag provides a mock function with given fields: opts, subject
func (_m *Flags) GetFlag(opts *bind.CallOpts, subject common.Address) (bool, error) {
	ret := _m.Called(opts, subject)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) bool); ok {
		r0 = rf(opts, subject)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlags provides a mock function with given fields: opts, subjects
func (_m *Flags) GetFlags(opts *bind.CallOpts, subjects []common.Address) ([]bool, error) {
	ret := _m.Called(opts, subjects)

	var r0 []bool
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, []common.Address) []bool); ok {
		r0 = rf(opts, subjects)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, []common.Address) error); ok {
		r1 = rf(opts, subjects)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAccess provides a mock function with given fields: opts, _user, _calldata
func (_m *Flags) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	ret := _m.Called(opts, _user, _calldata)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address, []byte) bool); ok {
		r0 = rf(opts, _user, _calldata)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address, []byte) error); ok {
		r1 = rf(opts, _user, _calldata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LowerFlags provides a mock function with given fields: opts, subjects
func (_m *Flags) LowerFlags(opts *bind.TransactOpts, subjects []common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, subjects)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address) *types.Transaction); ok {
		r0 = rf(opts, subjects)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address) error); ok {
		r1 = rf(opts, subjects)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *Flags) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseAddedAccess provides a mock function with given fields: log
func (_m *Flags) ParseAddedAccess(log types.Log) (*flags_wrapper.FlagsAddedAccess, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsAddedAccess
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsAddedAccess); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsAddedAccess)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseCheckAccessDisabled provides a mock function with given fields: log
func (_m *Flags) ParseCheckAccessDisabled(log types.Log) (*flags_wrapper.FlagsCheckAccessDisabled, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsCheckAccessDisabled
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsCheckAccessDisabled); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsCheckAccessDisabled)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseCheckAccessEnabled provides a mock function with given fields: log
func (_m *Flags) ParseCheckAccessEnabled(log types.Log) (*flags_wrapper.FlagsCheckAccessEnabled, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsCheckAccessEnabled
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsCheckAccessEnabled); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsCheckAccessEnabled)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFlagLowered provides a mock function with given fields: log
func (_m *Flags) ParseFlagLowered(log types.Log) (*flags_wrapper.FlagsFlagLowered, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsFlagLowered
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsFlagLowered); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsFlagLowered)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFlagRaised provides a mock function with given fields: log
func (_m *Flags) ParseFlagRaised(log types.Log) (*flags_wrapper.FlagsFlagRaised, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsFlagRaised
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsFlagRaised); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsFlagRaised)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *Flags) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	var r0 generated.AbigenLog
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferRequested provides a mock function with given fields: log
func (_m *Flags) ParseOwnershipTransferRequested(log types.Log) (*flags_wrapper.FlagsOwnershipTransferRequested, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsOwnershipTransferRequested
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsOwnershipTransferRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsOwnershipTransferRequested)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *Flags) ParseOwnershipTransferred(log types.Log) (*flags_wrapper.FlagsOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsOwnershipTransferred
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsOwnershipTransferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRaisingAccessControllerUpdated provides a mock function with given fields: log
func (_m *Flags) ParseRaisingAccessControllerUpdated(log types.Log) (*flags_wrapper.FlagsRaisingAccessControllerUpdated, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsRaisingAccessControllerUpdated
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsRaisingAccessControllerUpdated); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsRaisingAccessControllerUpdated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRemovedAccess provides a mock function with given fields: log
func (_m *Flags) ParseRemovedAccess(log types.Log) (*flags_wrapper.FlagsRemovedAccess, error) {
	ret := _m.Called(log)

	var r0 *flags_wrapper.FlagsRemovedAccess
	if rf, ok := ret.Get(0).(func(types.Log) *flags_wrapper.FlagsRemovedAccess); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flags_wrapper.FlagsRemovedAccess)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RaiseFlag provides a mock function with given fields: opts, subject
func (_m *Flags) RaiseFlag(opts *bind.TransactOpts, subject common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, subject)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RaiseFlags provides a mock function with given fields: opts, subjects
func (_m *Flags) RaiseFlags(opts *bind.TransactOpts, subjects []common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, subjects)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address) *types.Transaction); ok {
		r0 = rf(opts, subjects)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address) error); ok {
		r1 = rf(opts, subjects)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RaisingAccessController provides a mock function with given fields: opts
func (_m *Flags) RaisingAccessController(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAccess provides a mock function with given fields: opts, _user
func (_m *Flags) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _user)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetRaisingAccessController provides a mock function with given fields: opts, racAddress
func (_m *Flags) SetRaisingAccessController(opts *bind.TransactOpts, racAddress common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, racAddress)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, racAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, racAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, _to
func (_m *Flags) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _to)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchAddedAccess provides a mock function with given fields: opts, sink
func (_m *Flags) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsAddedAccess) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsAddedAccess) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsAddedAccess) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchCheckAccessDisabled provides a mock function with given fields: opts, sink
func (_m *Flags) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsCheckAccessDisabled) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsCheckAccessDisabled) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsCheckAccessDisabled) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchCheckAccessEnabled provides a mock function with given fields: opts, sink
func (_m *Flags) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsCheckAccessEnabled) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsCheckAccessEnabled) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsCheckAccessEnabled) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchFlagLowered provides a mock function with given fields: opts, sink, subject
func (_m *Flags) WatchFlagLowered(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsFlagLowered, subject []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, subject)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsFlagLowered, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsFlagLowered, []common.Address) error); ok {
		r1 = rf(opts, sink, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchFlagRaised provides a mock function with given fields: opts, sink, subject
func (_m *Flags) WatchFlagRaised(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsFlagRaised, subject []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, subject)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsFlagRaised, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsFlagRaised, []common.Address) error); ok {
		r1 = rf(opts, sink, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferRequested provides a mock function with given fields: opts, sink, from, to
func (_m *Flags) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsOwnershipTransferRequested, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsOwnershipTransferRequested, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, from, to
func (_m *Flags) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRaisingAccessControllerUpdated provides a mock function with given fields: opts, sink, previous, current
func (_m *Flags) WatchRaisingAccessControllerUpdated(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsRaisingAccessControllerUpdated, previous []common.Address, current []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, previous, current)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsRaisingAccessControllerUpdated, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, previous, current)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsRaisingAccessControllerUpdated, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, previous, current)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRemovedAccess provides a mock function with given fields: opts, sink
func (_m *Flags) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *flags_wrapper.FlagsRemovedAccess) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsRemovedAccess) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *flags_wrapper.FlagsRemovedAccess) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
