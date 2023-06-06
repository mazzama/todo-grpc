package repository

import (
	"context"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(ctx context.Context, item *entity.Item) error
	Update(ctx context.Context, item *entity.Item) (*entity.Item, error)
	FindOneByCriteria(ctx context.Context, criteria map[string]interface{}) (*entity.Item, error)
	FindManyByCriteria(ctx context.Context, criteria map[string]interface{}, orderBy ...string) ([]*entity.Item, error)
	Delete(ctx context.Context, itemID uint64) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		db: db,
	}
}

func (i itemRepository) Create(ctx context.Context, item *entity.Item) error {
	return i.db.WithContext(ctx).Create(&item).Error
}

func (i itemRepository) Update(ctx context.Context, item *entity.Item) (*entity.Item, error) {
	tx := i.db.WithContext(ctx).Save(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return item, nil
}

func (i itemRepository) FindOneByCriteria(ctx context.Context, criteria map[string]interface{}) (*entity.Item, error) {
	var item entity.Item

	q := i.db.WithContext(ctx).
		Where(criteria)

	res := q.First(&item)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &item, nil
}

func (i itemRepository) FindManyByCriteria(ctx context.Context, criteria map[string]interface{}, orderBy ...string) ([]*entity.Item, error) {
	var items []*entity.Item
	q := i.db.WithContext(ctx).
		Where(criteria)

	if len(orderBy) == 1 {
		q.Order(orderBy)
	}

	res := q.Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return items, nil
}

func (i itemRepository) Delete(ctx context.Context, itemID uint64) error {
	tx := i.db.WithContext(ctx).Where("id = ?", itemID).Delete(&entity.Item{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
