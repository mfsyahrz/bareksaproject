package repository

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
)

type TopicRepository interface {
	FindAll(ctx context.Context) ([]entity.Topic, error)
}
