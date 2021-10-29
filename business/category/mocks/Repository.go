// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	category "AltaStore/business/category"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteCategory provides a mock function with given fields: id, adminId
func (_m *Repository) DeleteCategory(id string, adminId string) error {
	ret := _m.Called(id, adminId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, adminId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCategoryByCode provides a mock function with given fields: code
func (_m *Repository) FindCategoryByCode(code string) (*category.Category, error) {
	ret := _m.Called(code)

	var r0 *category.Category
	if rf, ok := ret.Get(0).(func(string) *category.Category); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCategoryById provides a mock function with given fields: id
func (_m *Repository) FindCategoryById(id string) (*category.Category, error) {
	ret := _m.Called(id)

	var r0 *category.Category
	if rf, ok := ret.Get(0).(func(string) *category.Category); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
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

// GetAllCategory provides a mock function with given fields:
func (_m *Repository) GetAllCategory() (*[]category.Category, error) {
	ret := _m.Called()

	var r0 *[]category.Category
	if rf, ok := ret.Get(0).(func() *[]category.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]category.Category)
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

// InsertCategory provides a mock function with given fields: _a0
func (_m *Repository) InsertCategory(_a0 category.Category) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(category.Category) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCategory provides a mock function with given fields: id, _a1
func (_m *Repository) UpdateCategory(id string, _a1 category.Category) error {
	ret := _m.Called(id, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, category.Category) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
