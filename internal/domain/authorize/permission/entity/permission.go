package entity

import "time"

type Permission struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	GroupName   string     `json:"group_name"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedBy   *uint64    `json:"created_by"`
	UpdatedBy   *uint64    `json:"updated_by"`
	DeletedBy   *uint64    `json:"deleted_by"`
}
