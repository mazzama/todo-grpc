package entity

import (
	"github.com/mazzama/todo-grpc/internal/todo/constant"
	"time"
)

type Item struct {
	ID          uint64
	Name        string              `gorm:"type:varchar(50);not null"`
	Description string              `gorm:"type:varchar(255);not null"`
	Notes       string              `gorm:"type:varchar(100)"`
	Status      constant.ItemStatus `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time           `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time           `gorm:"default:CURRENT_TIMESTAMP"`
}
