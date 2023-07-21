package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Profile   string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
