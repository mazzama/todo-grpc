package it

import (
	"context"
	"fmt"
	"github.com/mazzama/todo-grpc/internal/todo/constant"
	"github.com/mazzama/todo-grpc/internal/todo/entity"
	"github.com/mazzama/todo-grpc/internal/todo/repository"
)

func (p *PostgresRepositoryTestSuite) TestPostgresItemRepositoryCreateTest() {
	newItem := &entity.Item{
		Name:        "Calculus",
		Description: "Integral",
		Notes:       "For tomorrow",
		Status:      constant.ItemStatusTodo,
	}

	r := repository.NewItemRepository(p.gormDB)
	p.Assert().NoError(r.Create(context.Background(), newItem))

	fmt.Println("Create")

	id := newItem.ID

	var result entity.Item
	p.Assert().NoError(p.gormDB.First(&result, entity.Item{ID: id}).Error)
	p.Assert().Equal(newItem.Name, result.Name)
	p.Assert().Equal(newItem.Description, result.Description)
	p.Assert().Equal(newItem.Notes, result.Notes)
	p.Assert().Equal(newItem.Status, result.Status)
}

func (p *PostgresRepositoryTestSuite) TestPostgresItemRepositoryUpdateTest() {
	newItem := &entity.Item{
		ID:          2,
		Name:        "Calculus",
		Description: "Integral",
		Notes:       "For tomorrow",
		Status:      constant.ItemStatusInProgress,
	}

	r := repository.NewItemRepository(p.gormDB)
	updatedItem, err := r.Update(context.Background(), newItem)
	p.Assert().NoError(err)

	p.Assert().Equal(newItem.Name, updatedItem.Name)
	p.Assert().Equal(newItem.Description, updatedItem.Description)
	p.Assert().Equal(newItem.Notes, updatedItem.Notes)
	p.Assert().Equal(newItem.Status, updatedItem.Status)
}

func (p *PostgresRepositoryTestSuite) TestPostgresItemRepositoryFindByIdTest() {
	r := repository.NewItemRepository(p.gormDB)

	newItem := &entity.Item{
		Name:        "Calculus",
		Description: "Integral",
		Notes:       "For tomorrow",
		Status:      constant.ItemStatusTodo,
	}

	err := r.Create(context.Background(), newItem)
	p.Assert().NoError(err)

	item, err := r.FindOneByCriteria(context.Background(), map[string]interface{}{
		"id": 1,
	})
	p.Assert().NoError(err)
	p.Assert().NotNil(item)
}

func (p *PostgresRepositoryTestSuite) TestPostgresItemRepositoryViewItemListTest() {
	r := repository.NewItemRepository(p.gormDB)

	var items = []*entity.Item{
		{
			Name:        "Calculus 1",
			Description: "Basic Number",
			Notes:       "For tomorrow",
			Status:      constant.ItemStatusTodo,
		},
		{
			Name:        "Calculus 2",
			Description: "Integral",
			Notes:       "For tomorrow",
			Status:      constant.ItemStatusTodo,
		},
	}

	for _, v := range items {
		err := r.Create(context.Background(), v)
		p.Assert().NoError(err)
	}

	items, err := r.FindManyByCriteria(context.Background(), map[string]interface{}{
		"status": constant.ItemStatusTodo,
	})
	p.Assert().NoError(err)
	p.Assert().Equal(2, len(items))
}
