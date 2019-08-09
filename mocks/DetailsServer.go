// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "github.com/sdgmf/go-project-sample/api/proto"

// DetailsServer is an autogenerated mock type for the DetailsServer type
type DetailsServer struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *DetailsServer) Get(_a0 context.Context, _a1 *proto.GetDetailRequest) (*proto.Detail, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.Detail
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetDetailRequest) *proto.Detail); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Detail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetDetailRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}