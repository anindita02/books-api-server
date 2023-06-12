// Code generated by mockery v2.15.0. DO NOT EDIT.

package store

import (
	models "github.com/anindita02/books-api-server/pkg/app/models"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: book
func (_m *Store) CreateBook(book *models.Book) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: id
func (_m *Store) DeleteBook(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBookByID provides a mock function with given fields: id
func (_m *Store) GetBookByID(id int) (*models.Book, error) {
	ret := _m.Called(id)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(int) *models.Book); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields:
func (_m *Store) GetBooks() ([]*models.Book, error) {
	ret := _m.Called()

	var r0 []*models.Book
	if rf, ok := ret.Get(0).(func() []*models.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Book)
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

// UpdateBook provides a mock function with given fields: book
func (_m *Store) UpdateBook(book *models.Book) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}