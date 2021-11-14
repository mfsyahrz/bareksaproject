package postgres

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cast"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
	"github.com/mfsyahrz/bareksaproject/internal/infrastructure/redis/cache"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type tagRepo struct {
	db    *sqlx.DB
	log   *log.Logger
	cache cache.Cache
}

func NewTagRepo(db *sqlx.DB, log *log.Logger, cache cache.Cache) repository.TagRepository {
	return &tagRepo{db, log, cache}
}

func (r *tagRepo) FindOne(ctx context.Context, id int64) (*entity.Tag, error) {
	var tag = &entity.Tag{}

	if err := r.cache.Get(ctx, createCacheKey(id), tag); err != nil {
		r.log.Error(err.Error())
	}

	if tag.ID != 0 {
		return tag, nil
	}

	var tagList []entity.Tag
	err := r.db.Select(&tagList, `SELECT * FROM tag WHERE id = $1`, id)
	if err != nil {
		r.log.Error(err.Error())
		return nil, constant.ErrDatabase
	}

	if len(tagList) == 0 {
		return nil, constant.ErrNotFound
	}

	if err = r.cache.Set(ctx, createCacheKey(id), tagList[0]); err != nil {
		r.log.Error(err.Error())
	}
	return &tagList[0], nil
}

func (r *tagRepo) FindAll(ctx context.Context) ([]entity.Tag, error) {
	var entities []entity.Tag
	err := r.db.Select(&entities, `SELECT * FROM tag`)
	if err != nil {
		r.log.Error(err.Error())
		return nil, constant.ErrDatabase
	}

	return entities, nil
}

func (r *tagRepo) Create(ctx context.Context, tag entity.Tag) error {
	var err error

	res, err := r.db.Exec(`insert into tag (name) values ($1)`, tag.Name)
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}

	rows, err := res.RowsAffected()
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}

	if rows != 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *tagRepo) UpdateByID(ctx context.Context, tag entity.Tag) error {

	res, err := r.db.Exec(`UPDATE tag SET name = $1,
		updated_date = now() WHERE id = $2`, tag.Name, tag.ID)
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return constant.ErrDatabase
	}

	if rows != 1 {
		return errors.New("no rows affected")
	}

	if err = r.cache.Set(ctx, createCacheKey(tag.ID), tag); err != nil {
		r.log.Error(err.Error())
	}

	return nil
}

func (r *tagRepo) DeleteByID(ctx context.Context, id int64) error {
	var err error

	res, err := r.db.Exec(`DELETE from tag WHERE id = $1`, id)
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}

	rows, err := res.RowsAffected()
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}
	if rows != 1 {
		return errors.New("no rows affected")
	}

	err = r.cache.Del(ctx, createCacheKey(id))
	if err != nil {
		r.log.Error(err.Error())
	}

	return nil
}

func createCacheKey(id int64) string {
	return "tag" + "_" + cast.ToString(id)
}
