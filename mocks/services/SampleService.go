// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "ddd-boilerplate/models"

	mock "github.com/stretchr/testify/mock"
)

// SampleService is an autogenerated mock type for the SampleService type
type SampleService struct {
	mock.Mock
}

// FindSampleByID provides a mock function with given fields: ctx, id
func (_m *SampleService) FindSampleByID(ctx context.Context, id int64) (*models.SampleAPIResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.SampleAPIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.SampleAPIResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.SampleAPIResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SampleAPIResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSampleService creates a new instance of SampleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSampleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SampleService {
	mock := &SampleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
