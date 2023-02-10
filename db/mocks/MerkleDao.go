// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import util "github.com/kyokan/plasma/util"

// MerkleDao is an autogenerated mock type for the MerkleDao type
type MerkleDao struct {
	mock.Mock
}

// Save provides a mock function with given fields: n
func (_m *MerkleDao) Save(n *util.MerkleNode) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(*util.MerkleNode) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveMany provides a mock function with given fields: ns
func (_m *MerkleDao) SaveMany(ns []util.MerkleNode) error {
	ret := _m.Called(ns)

	var r0 error
	if rf, ok := ret.Get(0).(func([]util.MerkleNode) error); ok {
		r0 = rf(ns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
