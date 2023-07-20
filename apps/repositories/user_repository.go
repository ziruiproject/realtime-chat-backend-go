package repositories

import (
	"context"
	"database/sql"

	"github.com/ziruiproject/realtime-chat-backend-go/apps/models"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user models.User) models.User
	Update(ctx context.Context, tx *sql.Tx, user models.User) models.User
	Delete(ctx context.Context, tx *sql.Tx, user models.User)
	GetById(ctx context.Context, tx *sql.Tx, id string) (models.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []models.User
}
