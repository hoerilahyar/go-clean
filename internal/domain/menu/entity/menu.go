package entity

import (
	"database/sql"
	"time"
)

type Menu struct {
	ID        uint64
	ParentID  *uint64
	Name      string
	Slug      string
	Path      string
	Icon      *string
	SortOrder int
	IsVisible bool
	IsActive  bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime

	CreatedBy uint64
	UpdatedBy *uint64
	DeletedBy *uint64
}
