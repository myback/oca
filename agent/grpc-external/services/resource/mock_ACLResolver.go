// Code generated by mockery v2.20.0. DO NOT EDIT.

package resource

import (
	acl "github.com/myback/oca/acl"
	mock "github.com/stretchr/testify/mock"

	resolver "github.com/myback/oca/acl/resolver"
)

// MockACLResolver is an autogenerated mock type for the ACLResolver type
type MockACLResolver struct {
	mock.Mock
}

// ResolveTokenAndDefaultMeta provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockACLResolver) ResolveTokenAndDefaultMeta(_a0 string, _a1 *acl.EnterpriseMeta, _a2 *acl.AuthorizerContext) (resolver.Result, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 resolver.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) resolver.Result); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(resolver.Result)
	}

	if rf, ok := ret.Get(1).(func(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockACLResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockACLResolver creates a new instance of MockACLResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockACLResolver(t mockConstructorTestingTNewMockACLResolver) *MockACLResolver {
	mock := &MockACLResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
