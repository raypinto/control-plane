// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/control-plane/components/provisioner/internal/apperrors"
import context "context"
import mock "github.com/stretchr/testify/mock"

// TenantUpdater is an autogenerated mock type for the TenantUpdater type
type TenantUpdater struct {
	mock.Mock
}

// GetAndUpdateTenant provides a mock function with given fields: runtimeID, ctx
func (_m *TenantUpdater) GetAndUpdateTenant(runtimeID string, ctx context.Context) apperrors.AppError {
	ret := _m.Called(runtimeID, ctx)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, context.Context) apperrors.AppError); ok {
		r0 = rf(runtimeID, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// GetTenant provides a mock function with given fields: ctx
func (_m *TenantUpdater) GetTenant(ctx context.Context) (string, apperrors.AppError) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(context.Context) apperrors.AppError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
