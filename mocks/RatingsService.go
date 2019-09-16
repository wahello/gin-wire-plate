// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/keyeMyria/gin-wire-plate/internal/pkg/models"

// RatingsService is an autogenerated mock type for the RatingsService type
type RatingsService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ID
func (_m *RatingsService) Get(ID uint64) (*models.Rating, error) {
	ret := _m.Called(ID)

	var r0 *models.Rating
	if rf, ok := ret.Get(0).(func(uint64) *models.Rating); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rating)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
