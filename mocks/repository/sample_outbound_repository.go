// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock_repository

import (
	context "context"
	"ddd-boilerplate/repositories/outbound/sample"

	mock "github.com/stretchr/testify/mock"
)

// SampleOutboundRepository is an autogenerated mock type for the SampleOutboundRepository type
type SampleOutboundRepository struct {
	mock.Mock
}

func (_m *SampleOutboundRepository) PostSampleAPI(ctx context.Context, request sample.SamplePostAPIRequest) (*sample.SamplePostAPIResponse, error) {
	//TODO implement me
	panic("implement me")
}

// FindSampleAPI provides a mock function with given fields: ctx, id
func (_m *SampleOutboundRepository) FindSampleAPI(ctx context.Context, id int64) (*sample.SampleAPIResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *sample.SampleAPIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*sample.SampleAPIResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *sample.SampleAPIResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sample.SampleAPIResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSampleOutboundRepository creates a new instance of SampleOutboundRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSampleOutboundRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SampleOutboundRepository {
	mock := &SampleOutboundRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
