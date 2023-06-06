package service

import (
	"context"
	"fmt"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	"github.com/mazzama/todo-grpc/internal/todo/mocks"
	"github.com/mazzama/todo-grpc/internal/todo/model"
	custom_error "github.com/mazzama/todo-grpc/pkg/error"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

type fields struct {
	*mocks.ItemRepository
}

func TestType(t *testing.T) {
	mspx := map[string]interface{}{}
	fmt.Printf("%T", mspx)
}

func TestItemService_Create(t *testing.T) {
	type args struct {
		ctx     context.Context
		request model.CreateItemRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		on     func(*fields)
		assert func(*fields)
	}{
		{
			name: "Given valid request when create item should return no error",
			in: &args{
				ctx: context.Background(),
				request: model.CreateItemRequest{
					Name:        "Homework",
					Description: "Basic Programming",
					Notes:       "For this weekend",
				},
			},
			out: nil,
			on: func(f *fields) {
				f.ItemRepository.On("Create",
					mock.AnythingOfType("*context.emptyCtx"),
					mock.AnythingOfType("*entity.Item")).
					Return(nil)
			},
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "Create", 1)
			},
		},
		{
			name: "Given invalid request when create item should return empty field error",
			in: &args{
				ctx: context.Background(),
				request: model.CreateItemRequest{
					Name:        "Homework",
					Description: "",
					Notes:       "For next weekend",
				},
			},
			out: custom_error.ErrEmptyField,
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "Create", 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				&mocks.ItemRepository{},
			}

			service := NewItemService(f.ItemRepository)
			if tt.on != nil {
				tt.on(f)
			}

			_, err := service.Create(tt.in.ctx, tt.in.request)
			if err != nil {
				if err.Error() != tt.out.Error() {
					t.Errorf("got %v, want %v", err, tt.out)
				}
			}
			if tt.assert != nil {
				tt.assert(f)
			}
		})
	}
}

func TestItemService_Update(t *testing.T) {
	type args struct {
		ctx     context.Context
		request model.UpdateItemRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		on     func(*fields)
		assert func(*fields)
	}{
		{
			name: "Given valid request when update item should return no error",
			in: &args{
				ctx: context.Background(),
				request: model.UpdateItemRequest{
					ID:          1,
					Name:        "Homework",
					Description: "Basic Programming",
					Notes:       "For this weekend",
					Status:      "IN_PROGRESS",
				},
			},
			out: nil,
			on: func(f *fields) {
				f.ItemRepository.
					On("FindOneByCriteria",
						mock.Anything,
						mock.Anything).
					Return(&entity.Item{
						ID:          1,
						Name:        "Linear Algebra Work",
						Description: "Phase 1",
						Notes:       "For Mid Exam",
						Status:      "TODO",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil).
					On("Update",
						mock.AnythingOfType("*context.emptyCtx"),
						mock.AnythingOfType("*entity.Item")).
					Return(&entity.Item{
						ID:          1,
						Name:        "Linear Algebra Work",
						Description: "Phase 1",
						Notes:       "For Mid Exam",
						Status:      "TODO",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
			},
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindOneByCriteria", 1)
				f.ItemRepository.AssertNumberOfCalls(t, "Update", 1)
			},
		},
		{
			name: "Given invalid request when update item should return not found error",
			in: &args{
				ctx: context.Background(),
				request: model.UpdateItemRequest{
					ID:          1,
					Name:        "Homework",
					Description: "Basic Programming",
					Notes:       "For next weekend",
					Status:      "UNKNOWN_STATUS",
				},
			},
			on: func(f *fields) {
				f.ItemRepository.
					On("FindOneByCriteria",
						mock.Anything,
						mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			out: gorm.ErrRecordNotFound,
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindOneByCriteria", 1)
				f.ItemRepository.AssertNumberOfCalls(t, "Update", 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				&mocks.ItemRepository{},
			}

			service := NewItemService(f.ItemRepository)
			if tt.on != nil {
				tt.on(f)
			}

			_, err := service.Update(tt.in.ctx, tt.in.request)
			if err != nil {
				if err.Error() != tt.out.Error() {
					t.Errorf("got %v, want %v", err, tt.out)
				}
			}
			if tt.assert != nil {
				tt.assert(f)
			}
		})
	}
}

func TestItemService_Delete(t *testing.T) {
	type args struct {
		ctx     context.Context
		request uint64
	}

	tests := []struct {
		name string
		in   *args
		out  error

		on     func(*fields)
		assert func(*fields)
	}{
		{
			name: "Given valid request when delete item should return empty field error",
			in: &args{
				ctx:     context.Background(),
				request: 1,
			},
			on: func(f *fields) {
				f.ItemRepository.
					On("Delete",
						mock.AnythingOfType("*context.emptyCtx"),
						mock.AnythingOfType("uint64")).
					Return(nil)
			},
			out: nil,
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "Delete", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				&mocks.ItemRepository{},
			}

			service := NewItemService(f.ItemRepository)
			if tt.on != nil {
				tt.on(f)
			}

			err := service.Delete(tt.in.ctx, tt.in.request)
			if err != nil {
				if err.Error() != tt.out.Error() {
					t.Errorf("got %v, want %v", err, tt.out)
				}
			}
			if tt.assert != nil {
				tt.assert(f)
			}
		})
	}
}

func TestItemService_FindItemById(t *testing.T) {
	type args struct {
		ctx     context.Context
		request uint64
	}

	tests := []struct {
		name string
		in   *args
		out  error

		on     func(*fields)
		assert func(*fields)
	}{
		{
			name: "Given valid request when find item by id should return no error",
			in: &args{
				ctx:     context.Background(),
				request: 100,
			},
			out: nil,
			on: func(f *fields) {
				f.ItemRepository.
					On("FindOneByCriteria",
						mock.Anything,
						mock.Anything).
					Return(&entity.Item{
						ID:          100,
						Name:        "Linear Algebra Work",
						Description: "Phase 1",
						Notes:       "For Mid Exam",
						Status:      "TODO",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}, nil)
			},
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindOneByCriteria", 1)
			},
		},
		{
			name: "Given invalid request when find item by id should return not found error",
			in: &args{
				ctx:     context.Background(),
				request: 1200,
			},
			on: func(f *fields) {
				f.ItemRepository.
					On("FindOneByCriteria",
						mock.Anything,
						mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			out: gorm.ErrRecordNotFound,
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindOneByCriteria", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				&mocks.ItemRepository{},
			}

			service := NewItemService(f.ItemRepository)
			if tt.on != nil {
				tt.on(f)
			}

			_, err := service.FindItemById(tt.in.ctx, tt.in.request)
			if err != nil {
				if err.Error() != tt.out.Error() {
					t.Errorf("got %v, want %v", err, tt.out)
				}
			}
			if tt.assert != nil {
				tt.assert(f)
			}
		})
	}
}

func TestItemService_ViewItemList(t *testing.T) {
	type args struct {
		ctx     context.Context
		request model.ViewItemListRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		on     func(*fields)
		assert func(*fields)
	}{
		{
			name: "Given valid request when find item by id should return no error",
			in: &args{
				ctx:     context.Background(),
				request: model.ViewItemListRequest{Status: "TODO"},
			},
			out: nil,
			on: func(f *fields) {
				f.ItemRepository.
					On("FindManyByCriteria",
						mock.Anything,
						mock.Anything).
					Return([]*entity.Item{
						{
							ID:          100,
							Name:        "Linear Algebra Work",
							Description: "Phase 1",
							Notes:       "For Mid Exam",
							Status:      "TODO",
							CreatedAt:   time.Now(),
							UpdatedAt:   time.Now(),
						},
					}, nil)
			},
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindManyByCriteria", 1)
			},
		},
		{
			name: "Given invalid request when find item by id should return not found error",
			in: &args{
				ctx:     context.Background(),
				request: model.ViewItemListRequest{Status: "INVALID"},
			},
			on: func(f *fields) {
				f.ItemRepository.
					On("FindManyByCriteria",
						mock.Anything,
						mock.Anything).
					Return(nil, gorm.ErrRecordNotFound)
			},
			out: gorm.ErrRecordNotFound,
			assert: func(f *fields) {
				f.ItemRepository.AssertNumberOfCalls(t, "FindManyByCriteria", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				&mocks.ItemRepository{},
			}

			service := NewItemService(f.ItemRepository)
			if tt.on != nil {
				tt.on(f)
			}

			_, err := service.ViewItemList(tt.in.ctx, tt.in.request)
			if err != nil {
				if err.Error() != tt.out.Error() {
					t.Errorf("got %v, want %v", err, tt.out)
				}
			}
			if tt.assert != nil {
				tt.assert(f)
			}
		})
	}
}
