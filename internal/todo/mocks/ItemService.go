// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mazzama/todo-grpc/internal/todo/model"
	mock "github.com/stretchr/testify/mock"
)

// ItemService is an autogenerated mock type for the ItemService type
type ItemService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, req
func (_m *ItemService) Create(ctx context.Context, req model.CreateItemRequest) (*model.CreateItemResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.CreateItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateItemRequest) (*model.CreateItemResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateItemRequest) *model.CreateItemResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CreateItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateItemRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, itemID
func (_m *ItemService) Delete(ctx context.Context, itemID uint64) error {
	ret := _m.Called(ctx, itemID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindItemById provides a mock function with given fields: ctx, itemID
func (_m *ItemService) FindItemById(ctx context.Context, itemID uint64) (*model.Item, error) {
	ret := _m.Called(ctx, itemID)

	var r0 *model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*model.Item, error)); ok {
		return rf(ctx, itemID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *model.Item); ok {
		r0 = rf(ctx, itemID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, itemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, req
func (_m *ItemService) Update(ctx context.Context, req model.UpdateItemRequest) (*model.Item, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateItemRequest) (*model.Item, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateItemRequest) *model.Item); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateItemRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ViewItemList provides a mock function with given fields: ctx, req
func (_m *ItemService) ViewItemList(ctx context.Context, req model.ViewItemListRequest) (*model.ViewItemListResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.ViewItemListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ViewItemListRequest) (*model.ViewItemListResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.ViewItemListRequest) *model.ViewItemListResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ViewItemListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.ViewItemListRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewItemService interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemService creates a new instance of ItemService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemService(t mockConstructorTestingTNewItemService) *ItemService {
	mock := &ItemService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}