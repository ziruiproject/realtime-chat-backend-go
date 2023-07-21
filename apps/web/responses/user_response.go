package responses

import "github.com/google/uuid"

type UserResponse struct {
	Id        uuid.UUID     `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Profile   string `json:"profile"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
