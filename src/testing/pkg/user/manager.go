// Code generated by mockery v2.1.0. DO NOT EDIT.

package user

import (
	context "context"

	models "github.com/goharbor/harbor/src/common/models"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"

	usermodels "github.com/goharbor/harbor/src/pkg/user/models"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *Manager) Get(ctx context.Context, id int) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, username
func (_m *Manager) GetByName(ctx context.Context, username string) (*models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) (usermodels.Users, error) {
	ret := _m.Called(ctx, query)

	var r0 usermodels.Users
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) usermodels.Users); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(usermodels.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
