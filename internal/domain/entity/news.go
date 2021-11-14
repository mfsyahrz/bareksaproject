package entity

import (
	"encoding/json"
	"time"

	"github.com/mfsyahrz/bareksaproject/internal/shared/constant"
	"github.com/spf13/cast"
)

type News struct {
	ID        int64               `db:"id" json:"id"`
	TopicID   int64               `db:"topic_id" json:"topic_id" validate:"required"`
	Title     string              `db:"title" json:"title" validate:"required,max=255"`
	Status    constant.NewsStatus `db:"status" json:"status" validate:"required,news_status"`
	Tag       []Tag               `db:"-" json:"tags"`
	CreatedAt time.Time           `db:"created_date" json:"-"`
	UpdatedAt time.Time           `db:"updated_date" json:"-"`
}

func (t News) GetKey(id int64) string {
	return "news" + "_" + cast.ToString(id)
}

func (t News) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *News) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

type NewsTagAggregate struct {
	ID        int64               `db:"id" json:"id"`
	TopicID   int64               `db:"topic_id"`
	Title     string              `db:"title"`
	Status    constant.NewsStatus `db:"status"`
	TagID     int64               `db:"tag_id"`
	TagName   string              `db:"tag_name"`
	CreatedAt time.Time           `db:"created_date" json:"created_at"`
	UpdatedAt time.Time           `db:"updated_date" json:"updated_at"`
}
