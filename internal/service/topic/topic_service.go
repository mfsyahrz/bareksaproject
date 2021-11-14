package service

import (
	"context"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
)

type TopicService interface {
	FindAll(ctx context.Context) ([]entity.Topic, error)
}

type topicService struct {
	topicRepo repository.TopicRepository
}

func NewTopicService(topicRepo repository.TopicRepository) TopicService {
	return &topicService{topicRepo}
}

func (s *topicService) FindAll(ctx context.Context) ([]entity.Topic, error) {
	return s.topicRepo.FindAll(ctx)
}
