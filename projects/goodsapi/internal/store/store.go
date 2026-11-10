package store

import (
	"context"
	"errors"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
)

var ErrNotFound = errors.New("good not found")

// getGood возвращает ErrNotFound
type GoodStore interface {
	DeleteGood(ctx context.Context, id int) error
	UpdateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (model.Good, error)
	GetGood(ctx context.Context, id int) (model.Good, error)
	CreateGood(ctx context.Context, name string, prise float64, CategoryID *int) (model.Good, error)
	ListGoods(ctx context.Context) ([]model.Good, error)
	PurchaseGood(ctx context.Context, idShop int, idGood int, count int) error
}
