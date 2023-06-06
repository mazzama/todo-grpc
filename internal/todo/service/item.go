package service

import (
	"context"
	"github.com/mazzama/todo-grpc/internal/todo/constant"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	"github.com/mazzama/todo-grpc/internal/todo/model"
	"github.com/mazzama/todo-grpc/internal/todo/repository"
)

type ItemService interface {
	Create(ctx context.Context, req model.CreateItemRequest) (*model.CreateItemResponse, error)
	Update(ctx context.Context, req model.UpdateItemRequest) (*model.Item, error)
	FindItemById(ctx context.Context, itemID uint64) (*model.Item, error)
	ViewItemList(ctx context.Context, req model.ViewItemListRequest) (*model.ViewItemListResponse, error)
	Delete(ctx context.Context, itemID uint64) error
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return &itemService{
		itemRepository: itemRepository,
	}
}

func (i itemService) Create(ctx context.Context, req model.CreateItemRequest) (*model.CreateItemResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	var item = entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Notes:       req.Description,
		Status:      constant.ItemStatusTodo,
	}
	err = i.itemRepository.Create(ctx, &item)
	if err != nil {
		return nil, err
	}

	return &model.CreateItemResponse{
		ID: item.ID,
	}, nil
}

func (i itemService) Update(ctx context.Context, req model.UpdateItemRequest) (*model.Item, error) {
	item, err := i.itemRepository.FindOneByCriteria(ctx, map[string]interface{}{
		"id": req.ID,
	})
	if err != nil {
		return nil, err
	}

	err = req.ValidateAndCompare(item)
	if err != nil {
		return nil, err
	}

	updatedItem, err := i.itemRepository.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return &model.Item{
		ID:          updatedItem.ID,
		Name:        updatedItem.Name,
		Description: updatedItem.Description,
		Notes:       updatedItem.Notes,
		Status:      updatedItem.Status.ToString(),
		CreatedAt:   updatedItem.CreatedAt,
		UpdatedAt:   updatedItem.UpdatedAt,
	}, nil
}

func (i itemService) FindItemById(ctx context.Context, itemID uint64) (*model.Item, error) {
	item, err := i.itemRepository.FindOneByCriteria(ctx, map[string]interface{}{
		"id": itemID,
	})
	if err != nil {
		return nil, err
	}

	return &model.Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Notes:       item.Notes,
		Status:      item.Status.ToString(),
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func (i itemService) ViewItemList(ctx context.Context, req model.ViewItemListRequest) (*model.ViewItemListResponse, error) {
	criteria := map[string]interface{}{}
	if req.Status != "" {
		criteria["status"] = req.Status
	}

	items, err := i.itemRepository.FindManyByCriteria(ctx, criteria)
	if err != nil {
		return nil, err
	}
	var itemResponse []model.Item
	for _, item := range items {
		var resp = model.Item{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Notes:       item.Notes,
			Status:      item.Status.ToString(),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}

		itemResponse = append(itemResponse, resp)
	}

	return &model.ViewItemListResponse{
		Total: len(itemResponse),
		Items: itemResponse,
	}, nil
}

func (i itemService) Delete(ctx context.Context, itemID uint64) error {
	return i.itemRepository.Delete(ctx, itemID)
}
