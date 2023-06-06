package model

import (
	"github.com/mazzama/todo-grpc/internal/todo/constant"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	custom_error "github.com/mazzama/todo-grpc/pkg/error"
	"time"
)

type CreateItemRequest struct {
	Name        string
	Description string
	Notes       string
}

func (r CreateItemRequest) Validate() error {
	if len(r.Name) < 1 || len(r.Description) < 1 {
		return custom_error.ErrEmptyField
	}

	return nil
}

type CreateItemResponse struct {
	ID uint64
}

type UpdateItemRequest struct {
	ID          uint64
	Name        string
	Description string
	Notes       string
	Status      string
}

func (r UpdateItemRequest) ValidateAndCompare(item *entity.Item) error {
	if len(r.Name) < 1 || len(r.Description) < 1 {
		return custom_error.ErrEmptyField
	}

	if r.Name != item.Name {
		item.Name = r.Name
	}

	if r.Description != item.Description {
		item.Description = r.Description
	}

	if r.Notes != item.Notes {
		item.Notes = r.Notes
	}

	status, ok := constant.MapItemStatus[r.Status]
	if !ok {
		return custom_error.ErrInvalidStatus
	}

	if status != item.Status {
		item.Status = status
	}

	return nil
}

type Item struct {
	ID          uint64
	Name        string
	Description string
	Notes       string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ViewItemListRequest struct {
	Status string
}

type ViewItemListResponse struct {
	Total int
	Items []Item
}
