package constant

import "errors"

type NewsStatus string

func (r NewsStatus) String() string {
	return string(r)
}

const (
	TypeNewsStatusDraft     = NewsStatus("draft")
	TypeNewsStatusDeleted   = NewsStatus("deleted")
	TypeNewsStatusPublished = NewsStatus("publish")
)

const (
	ResponseSuccess = "Success"
	IDEmpty         = "ID Cannot Be Empty"
)

var (
	ErrDatabase = errors.New("database error")
	ErrNotFound = errors.New("data not found")
)
