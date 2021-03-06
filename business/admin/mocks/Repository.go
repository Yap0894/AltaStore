// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admin "AltaStore/business/admin"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteAdmin provides a mock function with given fields: Admin
func (_m *Repository) DeleteAdmin(Admin admin.Admin) error {
	ret := _m.Called(Admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.Admin) error); ok {
		r0 = rf(Admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAdminByEmail provides a mock function with given fields: email
func (_m *Repository) FindAdminByEmail(email string) (*admin.Admin, error) {
	ret := _m.Called(email)

	var r0 *admin.Admin
	if rf, ok := ret.Get(0).(func(string) *admin.Admin); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAdminByEmailAndPassword provides a mock function with given fields: email, password
func (_m *Repository) FindAdminByEmailAndPassword(email string, password string) (*admin.Admin, error) {
	ret := _m.Called(email, password)

	var r0 *admin.Admin
	if rf, ok := ret.Get(0).(func(string, string) *admin.Admin); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAdminByID provides a mock function with given fields: id
func (_m *Repository) FindAdminByID(id string) (*admin.Admin, error) {
	ret := _m.Called(id)

	var r0 *admin.Admin
	if rf, ok := ret.Get(0).(func(string) *admin.Admin); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAdmin provides a mock function with given fields: Admin
func (_m *Repository) InsertAdmin(Admin admin.Admin) error {
	ret := _m.Called(Admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.Admin) error); ok {
		r0 = rf(Admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdmin provides a mock function with given fields: Admin
func (_m *Repository) UpdateAdmin(Admin admin.Admin) error {
	ret := _m.Called(Admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.Admin) error); ok {
		r0 = rf(Admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdminPassword provides a mock function with given fields: Admin
func (_m *Repository) UpdateAdminPassword(Admin admin.Admin) error {
	ret := _m.Called(Admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.Admin) error); ok {
		r0 = rf(Admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
