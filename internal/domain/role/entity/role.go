package entity

import "time"

type Role struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description *string    `json:"description"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedBy   *uint64    `json:"created_by"`
	UpdatedBy   *uint64    `json:"updated_by"`
	DeletedBy   *uint64    `json:"deleted_by"`
}
