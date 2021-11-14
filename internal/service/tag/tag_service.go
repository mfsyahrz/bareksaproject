package service

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
)

type TagService interface {
	FindAll(ctx context.Context) ([]entity.Tag, error)
	FindOne(ctx context.Context, id int64) (*entity.Tag, error)
	Create(ctx context.Context, tag entity.Tag) error
	UpdateByID(ctx context.Context, tag entity.Tag) error
	DeleteByID(ctx context.Context, id int64) error
}

type tagService struct {
	tagRepo repository.TagRepository
}

func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{tagRepo}
}

func (s *tagService) FindAll(ctx context.Context) ([]entity.Tag, error) {
	return s.tagRepo.FindAll(ctx)
}

func (s *tagService) FindOne(ctx context.Context, tagID int64) (*entity.Tag, error) {
	return s.tagRepo.FindOne(ctx, tagID)
}

func (s *tagService) Create(ctx context.Context, news entity.Tag) error {
	return s.tagRepo.Create(ctx, news)
}

func (s *tagService) UpdateByID(ctx context.Context, tag entity.Tag) error {
	return s.tagRepo.UpdateByID(ctx, tag)
}

func (s *tagService) DeleteByID(ctx context.Context, tagID int64) error {
	return s.tagRepo.DeleteByID(ctx, tagID)
}
