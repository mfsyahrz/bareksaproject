package repository

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
)

type TagRepository interface {
	FindAll(ctx context.Context) ([]entity.Tag, error)
	FindOne(ctx context.Context, id int64) (*entity.Tag, error)
	Create(ctx context.Context, tag entity.Tag) error
	UpdateByID(ctx context.Context, tag entity.Tag) error
	DeleteByID(ctx context.Context, id int64) error
}
