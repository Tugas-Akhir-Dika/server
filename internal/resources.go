package internal

import (
	"SDUI_Server/internal/model/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MainResources struct {
	Db *sqlx.DB
}

func NewRsc(cfg *config.Config) (*MainResources, error) {
	db, err := sqlx.Open("postgres", cfg.Database)
	if err != nil {
		return nil, err
	}
	return &MainResources{
		Db: db,
	}, nil
}
