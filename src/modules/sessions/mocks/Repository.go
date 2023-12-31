// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	model "NoteApp/src/model"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: session
func (_m *Repository) Create(session model.Session) (model.Session, error) {
	ret := _m.Called(session)

	var r0 model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Session) (model.Session, error)); ok {
		return rf(session)
	}
	if rf, ok := ret.Get(0).(func(model.Session) model.Session); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Get(0).(model.Session)
	}

	if rf, ok := ret.Get(1).(func(model.Session) error); ok {
		r1 = rf(session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: sessionID
func (_m *Repository) Delete(sessionID string) (model.Session, error) {
	ret := _m.Called(sessionID)

	var r0 model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Session, error)); ok {
		return rf(sessionID)
	}
	if rf, ok := ret.Get(0).(func(string) model.Session); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Get(0).(model.Session)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: sessionID
func (_m *Repository) Find(sessionID string) (model.Session, error) {
	ret := _m.Called(sessionID)

	var r0 model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Session, error)); ok {
		return rf(sessionID)
	}
	if rf, ok := ret.Get(0).(func(string) model.Session); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Get(0).(model.Session)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUser provides a mock function with given fields: userID
func (_m *Repository) FindByUser(userID uint) (model.Session, error) {
	ret := _m.Called(userID)

	var r0 model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (model.Session, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) model.Session); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(model.Session)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
