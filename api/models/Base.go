package models

import "time"

type Base struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
