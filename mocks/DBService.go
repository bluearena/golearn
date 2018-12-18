// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import golearn "github.com/sergeiten/golearn"
import mock "github.com/stretchr/testify/mock"

// DBService is an autogenerated mock type for the DBService type
type DBService struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *DBService) Close() {
	_m.Called()
}

// ExistUser provides a mock function with given fields: user
func (_m *DBService) ExistUser(user golearn.User) (bool, error) {
	ret := _m.Called(user)

	var r0 bool
	if rf, ok := ret.Get(0).(func(golearn.User) bool); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(golearn.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: _a0
func (_m *DBService) GetState(_a0 string) (golearn.State, error) {
	ret := _m.Called(_a0)

	var r0 golearn.State
	if rf, ok := ret.Get(0).(func(string) golearn.State); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(golearn.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userid
func (_m *DBService) GetUser(userid string) (golearn.User, error) {
	ret := _m.Called(userid)

	var r0 golearn.User
	if rf, ok := ret.Get(0).(func(string) golearn.User); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Get(0).(golearn.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: user
func (_m *DBService) InsertUser(user golearn.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(golearn.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertWord provides a mock function with given fields: _a0
func (_m *DBService) InsertWord(_a0 golearn.Row) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(golearn.Row) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RandomAnswers provides a mock function with given fields: q, limit
func (_m *DBService) RandomAnswers(q golearn.Row, limit int) ([]golearn.Row, error) {
	ret := _m.Called(q, limit)

	var r0 []golearn.Row
	if rf, ok := ret.Get(0).(func(golearn.Row, int) []golearn.Row); ok {
		r0 = rf(q, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]golearn.Row)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(golearn.Row, int) error); ok {
		r1 = rf(q, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandomQuestion provides a mock function with given fields:
func (_m *DBService) RandomQuestion() (golearn.Row, error) {
	ret := _m.Called()

	var r0 golearn.Row
	if rf, ok := ret.Get(0).(func() golearn.Row); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(golearn.Row)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetState provides a mock function with given fields: _a0
func (_m *DBService) ResetState(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetState provides a mock function with given fields: _a0
func (_m *DBService) SetState(_a0 golearn.State) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(golearn.State) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUserMode provides a mock function with given fields: userID, mode
func (_m *DBService) SetUserMode(userID string, mode string) error {
	ret := _m.Called(userID, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: user
func (_m *DBService) UpdateUser(user golearn.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(golearn.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
