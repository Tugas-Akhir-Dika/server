package item

import (
	"SDUI_Server/internal/model/entity"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
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

func (r *Repository) AddItem() error {
	return nil
}

func (r *Repository) GetItemById(ctx context.Context, id int64) (*entity.ItemEntity, error) {
	item := &entity.ItemEntity{}
	err := r.db.GetContext(ctx, item, "SELECT id, title,price, description, category, image, rate, count FROM items WHERE id = $1 LIMIT 1", id)
	if err != nil {
		log.Error(err)
		return nil, errors.New("cannot find item")
	}
	return item, nil
}

func (r *Repository) GetItems(ctx context.Context, page uint64, limit uint64) ([]entity.ItemEntity, error) {
	offset := (page - 1) * limit
	query := "SELECT id, title,price, description, category, image, rate, count FROM items ORDER BY id ASC LIMIT $1 OFFSET $2"
	items := []entity.ItemEntity{}
	err := r.db.SelectContext(ctx, &items, query, limit, offset)
	if err != nil {
		log.Error(err)
		return items, errors.New("cannot find item")
	}
	return items, nil
}
