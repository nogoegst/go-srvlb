// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import srv "github.com/improbable-eng/go-srvlb/srv"

// Resolver is an autogenerated mock type for the Resolver type
type Resolver struct {
	mock.Mock
}

// Lookup provides a mock function with given fields: domainName
func (_m *Resolver) Lookup(domainName string) ([]*srv.Target, error) {
	ret := _m.Called(domainName)

	var r0 []*srv.Target
	if rf, ok := ret.Get(0).(func(string) []*srv.Target); ok {
		r0 = rf(domainName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*srv.Target)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(domainName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
