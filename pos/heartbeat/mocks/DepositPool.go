// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import merkle "github.com/multivactech/MultiVAC/model/merkle"
import mock "github.com/stretchr/testify/mock"
import multivacaddress "github.com/multivactech/MultiVAC/model/chaincfg/multivacaddress"
import shard "github.com/multivactech/MultiVAC/model/shard"
import state "github.com/multivactech/MultiVAC/processor/shared/state"
import wire "github.com/multivactech/MultiVAC/model/wire"

// DepositPool is an autogenerated mock type for the DepositPool type
type DepositPool struct {
	mock.Mock
}

// Add provides a mock function with given fields: out, height, proof
func (_m *DepositPool) Add(out *wire.OutPoint, height wire.BlockHeight, proof *merkle.MerklePath) error {
	ret := _m.Called(out, height, proof)

	var r0 error
	if rf, ok := ret.Get(0).(func(*wire.OutPoint, wire.BlockHeight, *merkle.MerklePath) error); ok {
		r0 = rf(out, height, proof)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: _a0, address
func (_m *DepositPool) GetAll(_a0 shard.Index, address multivacaddress.Address) ([]*wire.OutWithProof, error) {
	ret := _m.Called(_a0, address)

	var r0 []*wire.OutWithProof
	if rf, ok := ret.Get(0).(func(shard.Index, multivacaddress.Address) []*wire.OutWithProof); ok {
		r0 = rf(_a0, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*wire.OutWithProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(shard.Index, multivacaddress.Address) error); ok {
		r1 = rf(_a0, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBiggest provides a mock function with given fields: address
func (_m *DepositPool) GetBiggest(address multivacaddress.Address) ([]byte, error) {
	ret := _m.Called(address)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(multivacaddress.Address) []byte); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(multivacaddress.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lock provides a mock function with given fields: out
func (_m *DepositPool) Lock(out *wire.OutPoint) error {
	ret := _m.Called(out)

	var r0 error
	if rf, ok := ret.Get(0).(func(*wire.OutPoint) error); ok {
		r0 = rf(out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: out
func (_m *DepositPool) Remove(out *wire.OutPoint) error {
	ret := _m.Called(out)

	var r0 error
	if rf, ok := ret.Get(0).(func(*wire.OutPoint) error); ok {
		r0 = rf(out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: update, _a1, height
func (_m *DepositPool) Update(update *state.Update, _a1 shard.Index, height wire.BlockHeight) error {
	ret := _m.Called(update, _a1, height)

	var r0 error
	if rf, ok := ret.Get(0).(func(*state.Update, shard.Index, wire.BlockHeight) error); ok {
		r0 = rf(update, _a1, height)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: address, proof
func (_m *DepositPool) Verify(address multivacaddress.Address, proof []byte) bool {
	ret := _m.Called(address, proof)

	var r0 bool
	if rf, ok := ret.Get(0).(func(multivacaddress.Address, []byte) bool); ok {
		r0 = rf(address, proof)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
