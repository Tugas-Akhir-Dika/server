package signin

import (
	"SDUI_Server/internal"
	"SDUI_Server/internal/model/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) SignInUser(ctx context.Context, username string, password string) (string, error) {
	return internal.AUTH_KEY, nil
}

func (r *Repository) GetUserInformation(ctx context.Context) entity.UserEntity {
	return entity.UserEntity{
		Username: "test123",
		Email:    "test@test.com",
		Photo:    "https://pbs.twimg.com/media/FyVXE2ZaIAANAgu?format=jpg&name=medium",
	}
}
