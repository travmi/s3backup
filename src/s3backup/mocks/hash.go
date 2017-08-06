package mocks

import mock "github.com/stretchr/testify/mock"

// Hash is an autogenerated mock type for the Hash type
type Hash struct {
	mock.Mock
}

// Calculate provides a mock function with given fields: filePath
func (_m *Hash) Calculate(filePath string) (string, error) {
	ret := _m.Called(filePath)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: filePath, expectedChecksum
func (_m *Hash) Verify(filePath string, expectedChecksum string) error {
	ret := _m.Called(filePath, expectedChecksum)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(filePath, expectedChecksum)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}