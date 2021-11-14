package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cast"

	"github.com/mfsyahrz/bareksaproject/internal/domain/entity"
	"github.com/mfsyahrz/bareksaproject/internal/domain/repository"
	"github.com/mfsyahrz/bareksaproject/internal/infrastructure/redis/cache"
	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
)

type newsRepo struct {
	db    *sqlx.DB
	log   *log.Logger
	cache cache.Cache
}

func NewNewsRepo(db *sqlx.DB, log *log.Logger, cache cache.Cache) repository.NewsRepository {
	return &newsRepo{db, log, cache}
}

func (r *newsRepo) FindOne(ctx context.Context, id int64) (*entity.News, error) {
	var news = &entity.News{}

	if err := r.cache.Get(ctx, createCacheKey(id), news); err != nil {
		r.log.Error(err.Error())
	}

	if news.ID != 0 {
		return news, nil
	}

	var newsList []entity.News
	err := r.db.Select(&newsList, `SELECT * FROM news WHERE id = $1`, id)
	if err != nil {
		r.log.Error(err.Error())
		return nil, constant.ErrDatabase
	}

	if len(newsList) == 0 {
		return nil, constant.ErrNotFound
	}

	if err = r.cache.Set(ctx, createCacheKey(id), newsList[0]); err != nil {
		r.log.Error(err.Error())
	}
	return &newsList[0], nil
}

func (r *newsRepo) FindAll(ctx context.Context, status constant.NewsStatus, topicID int64) ([]entity.News, error) {
	query := `SELECT n.id, n.title, n.topic_id, n.status "tag_name" FROM news n	
				JOIN topic tp on n.topic_id = tp.id
				WHERE 1 = 1`

	var args []interface{}
	sb := strings.Builder{}
	_, err := sb.WriteString(query)
	if err != nil {
		r.log.Error(err.Error())
		return nil, constant.ErrDatabase
	}

	if status != "" {
		args = append(args, status)
		sb.WriteString(" AND status = $" + cast.ToString(len(args)))
	}

	if topicID != 0 {
		args = append(args, topicID)
		sb.WriteString(" AND topic_id = $" + cast.ToString(len(args)))
	}

	rows, err := r.db.Query(sb.String(), args...)
	if err != nil {
		r.log.Error(err.Error())
		return nil, constant.ErrDatabase
	}

	var newsList = []entity.News{}
	for rows.Next() {
		news := entity.News{}

		err = rows.Scan(
			&news.ID,
			&news.Title,
			&news.TopicID,
			&news.Status,
		)
		if err != nil {
			r.log.Error(err.Error())
			return nil, constant.ErrDatabase
		}

		var newsTag []entity.NewsTagAggregate

		err := r.db.Select(&newsTag, `SELECT n.id, t.id as "tag_id", t.name as "tag_name" FROM news_tag nt join tag t on t.id = nt.tag_id
		 join news n on  n.id = nt.news_id where n.id = $1`, news.ID)
		if err != nil {
			r.log.Error(err.Error())
			return nil, constant.ErrDatabase
		}

		for _, val := range newsTag {
			news.Tag = append(news.Tag, entity.Tag{
				ID:   val.TagID,
				Name: val.TagName,
			})
		}

		newsList = append(newsList, news)
	}
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return newsList, nil
}

func (r *newsRepo) Create(ctx context.Context, news entity.News) error {

	var newsID int64
	err := r.db.QueryRowx(`insert into news (topic_id, title, status) 
	values ($1, $2, $3) RETURNING id`, news.TopicID, news.Title, news.Status).Scan(&newsID)
	if err != nil {
		r.log.Error(err.Error())
		return constant.ErrDatabase
	}

	for _, tag := range news.Tag {
		res, er := r.db.Exec(`insert into news_tag (news_id, tag_id) 
		values ($1, $2)`, newsID, tag.ID)
		if er != nil {
			err = er
			r.log.Error(err.Error())
			return constant.ErrDatabase
		}

		_, err = res.RowsAffected()
		if err != nil {
			r.log.Error(err.Error())
			return constant.ErrDatabase
		}
	}

	return nil
}

func (r *newsRepo) UpdateByID(ctx context.Context, news entity.News) error {
	res, err := r.db.Exec(`UPDATE news SET title = $1, status = $2,
	 updated_date = now() WHERE id = $3`, news.Title, news.Status, news.ID)
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

	if err = r.cache.Set(ctx, createCacheKey(news.ID), news); err != nil {
		r.log.Error(err.Error())
	}

	return nil

}

func (r *newsRepo) DeleteByID(ctx context.Context, id int64) error {
	res, err := r.db.Exec(`DELETE from news WHERE id = $1`, id)
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
	return "news" + "_" + cast.ToString(id)
}
