// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	types "github.com/ethereum/go-ethereum/core/types"
)

// KeyStoreInterface is an autogenerated mock type for the KeyStoreInterface type
type KeyStoreInterface struct {
	mock.Mock
}

// AddKey provides a mock function with given fields: key
func (_m *KeyStoreInterface) AddKey(key *models.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AllKeys provides a mock function with given fields:
func (_m *KeyStoreInterface) AllKeys() ([]models.Key, error) {
	ret := _m.Called()

	var r0 []models.Key
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewKey provides a mock function with given fields:
func (_m *KeyStoreInterface) CreateNewKey() (models.Key, error) {
	ret := _m.Called()

	var r0 models.Key
	if rf, ok := ret.Get(0).(func() models.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureFundingKey provides a mock function with given fields:
func (_m *KeyStoreInterface) EnsureFundingKey() (models.Key, bool, error) {
	ret := _m.Called()

	var r0 models.Key
	if rf, ok := ret.Get(0).(func() models.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ExportKey provides a mock function with given fields: address, newPassword
func (_m *KeyStoreInterface) ExportKey(address common.Address, newPassword string) ([]byte, error) {
	ret := _m.Called(address, newPassword)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Address, string) []byte); ok {
		r0 = rf(address, newPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, string) error); ok {
		r1 = rf(address, newPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FundingKeys provides a mock function with given fields:
func (_m *KeyStoreInterface) FundingKeys() ([]models.Key, error) {
	ret := _m.Called()

	var r0 []models.Key
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoundRobinAddress provides a mock function with given fields: addresses
func (_m *KeyStoreInterface) GetRoundRobinAddress(addresses ...common.Address) (common.Address, error) {
	_va := make([]interface{}, len(addresses))
	for _i := range addresses {
		_va[_i] = addresses[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(...common.Address) common.Address); ok {
		r0 = rf(addresses...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...common.Address) error); ok {
		r1 = rf(addresses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasDBSendingKeys provides a mock function with given fields:
func (_m *KeyStoreInterface) HasDBSendingKeys() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasSendingKeyWithAddress provides a mock function with given fields: address
func (_m *KeyStoreInterface) HasSendingKeyWithAddress(address common.Address) (bool, error) {
	ret := _m.Called(address)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportKey provides a mock function with given fields: keyJSON, oldPassword
func (_m *KeyStoreInterface) ImportKey(keyJSON []byte, oldPassword string) (models.Key, error) {
	ret := _m.Called(keyJSON, oldPassword)

	var r0 models.Key
	if rf, ok := ret.Get(0).(func([]byte, string) models.Key); ok {
		r0 = rf(keyJSON, oldPassword)
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, oldPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportKeyFileToDB provides a mock function with given fields: keyPath
func (_m *KeyStoreInterface) ImportKeyFileToDB(keyPath string) (models.Key, error) {
	ret := _m.Called(keyPath)

	var r0 models.Key
	if rf, ok := ret.Get(0).(func(string) models.Key); ok {
		r0 = rf(keyPath)
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyByAddress provides a mock function with given fields: address
func (_m *KeyStoreInterface) KeyByAddress(address common.Address) (models.Key, error) {
	ret := _m.Called(address)

	var r0 models.Key
	if rf, ok := ret.Get(0).(func(common.Address) models.Key); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveKey provides a mock function with given fields: address, hardDelete
func (_m *KeyStoreInterface) RemoveKey(address common.Address, hardDelete bool) (models.Key, error) {
	ret := _m.Called(address, hardDelete)

	var r0 models.Key
	if rf, ok := ret.Get(0).(func(common.Address, bool) models.Key); ok {
		r0 = rf(address, hardDelete)
	} else {
		r0 = ret.Get(0).(models.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, bool) error); ok {
		r1 = rf(address, hardDelete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendingKeys provides a mock function with given fields:
func (_m *KeyStoreInterface) SendingKeys() ([]models.Key, error) {
	ret := _m.Called()

	var r0 []models.Key
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTx provides a mock function with given fields: fromAddress, tx, chainID
func (_m *KeyStoreInterface) SignTx(fromAddress common.Address, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(fromAddress, tx, chainID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(common.Address, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(fromAddress, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Transaction, *big.Int) error); ok {
		r1 = rf(fromAddress, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields: password
func (_m *KeyStoreInterface) Unlock(password string) error {
	ret := _m.Called(password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
