package responses

import "github.com/google/uuid"

type AuthResponse struct {
	UserId  *uuid.UUID
	Token   string
	Success bool
}
