package service

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type NewsService interface {
	FindAll(ctx context.Context, status constant.NewsStatus, topicID int64) ([]entity.News, error)
	Create(ctx context.Context, news entity.News) error
	UpdateByID(ctx context.Context, news entity.News) error
	DeleteByID(ctx context.Context, id int64) error
	FindOne(ctx context.Context, id int64) (*entity.News, error)
}

type newsService struct {
	newsRepo repository.NewsRepository
}

func NewNewsService(newsRepo repository.NewsRepository) NewsService {
	return &newsService{newsRepo}
}

func (s *newsService) FindAll(ctx context.Context, status constant.NewsStatus, topicID int64) ([]entity.News, error) {
	return s.newsRepo.FindAll(ctx, status, topicID)
}

func (s *newsService) FindOne(ctx context.Context, newsID int64) (*entity.News, error) {
	return s.newsRepo.FindOne(ctx, newsID)
}

func (s *newsService) Create(ctx context.Context, news entity.News) error {
	return s.newsRepo.Create(ctx, news)
}

func (s *newsService) UpdateByID(ctx context.Context, news entity.News) error {
	return s.newsRepo.UpdateByID(ctx, news)
}

func (s *newsService) DeleteByID(ctx context.Context, newsID int64) error {
	return s.newsRepo.DeleteByID(ctx, newsID)
}
