package entity

import (
	"encoding/json"
	"time"
)

type Tag struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" validate:"required,max=255"`
	CreatedAt time.Time `db:"created_date" json:"-"`
	UpdatedAt time.Time `db:"updated_date" json:"-"`
}

func (t Tag) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Tag) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &t)
}
