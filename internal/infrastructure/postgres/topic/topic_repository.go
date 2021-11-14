package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
)

type topicRepo struct {
	db  *sqlx.DB
	log *log.Logger
}

func NewTopicRepo(db *sqlx.DB, log *log.Logger) repository.TopicRepository {
	return &topicRepo{db, log}
}

func (r *topicRepo) FindAll(ctx context.Context) ([]entity.Topic, error) {
	var entities []entity.Topic
	err := r.db.Select(&entities, `SELECT * FROM topic`)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return entities, nil
}
