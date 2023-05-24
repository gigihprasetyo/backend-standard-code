// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/pasarin-tech/pasarin-core/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// CityService is an autogenerated mock type for the CityService type
type CityService struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, params
func (_m *CityService) GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.CityTransformer, *domain.CursorTransform, error) {
	ret := _m.Called(ctx, params)

	var r0 []*domain.CityTransformer
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CityParams) []*domain.CityTransformer); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.CityTransformer)
		}
	}

	var r1 *domain.CursorTransform
	if rf, ok := ret.Get(1).(func(context.Context, *domain.CityParams) *domain.CursorTransform); ok {
		r1 = rf(ctx, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CursorTransform)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.CityParams) error); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetail provides a mock function with given fields: ctx, params
func (_m *CityService) GetDetail(ctx context.Context, params *domain.CityParams) (*domain.CityTransformer, error) {
	ret := _m.Called(ctx, params)

	var r0 *domain.CityTransformer
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CityParams) *domain.CityTransformer); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CityTransformer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.CityParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCityService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCityService creates a new instance of CityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCityService(t mockConstructorTestingTNewCityService) *CityService {
	mock := &CityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}