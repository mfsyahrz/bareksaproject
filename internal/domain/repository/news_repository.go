package repository

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type NewsRepository interface {
	FindAll(ctx context.Context, status constant.NewsStatus, topicID int64) ([]entity.News, error)
	FindOne(ctx context.Context, id int64) (*entity.News, error)
	Create(ctx context.Context, news entity.News) error
	UpdateByID(ctx context.Context, news entity.News) error
	DeleteByID(ctx context.Context, id int64) error
}
