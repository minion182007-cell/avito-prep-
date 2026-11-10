package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
)

type PgStore struct {
	db *sql.DB
}

func NewPgStore(db *sql.DB) *PgStore { return &PgStore{db: db} }
func (s *PgStore) PurchaseGood(ctx context.Context, idShop int, idGood int, count int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	var curentcount int
	err = tx.QueryRowContext(ctx, "SELECT count FROM Good_Shop where id_good=$1 AND id_shop=$2 FOR UPDATE ", idGood, idShop).Scan(&curentcount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return fmt.Errorf("chek stock: %w", err)
	}
	if curentcount < count {
		return errors.New("not enough stock")
	}
	_, err = tx.ExecContext(ctx, "UPDATE Good_Shop SET count=count-$1 where id_good=$2 AND id_shop=$3", count, idGood, idShop)
	if err != nil {
		return fmt.Errorf("update stock: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}
func (s *PgStore) DeleteGood(ctx context.Context, id int) error {
	res, err := s.db.ExecContext(ctx, "DELETE from good where id=$1", id)
	if err != nil {
		return fmt.Errorf("delete good:%d %w", id, err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete good:%d %w", id, err)
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
func (s *PgStore) UpdateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (model.Good, error) {
	var g model.Good
	err := s.db.QueryRowContext(ctx, "UPDATE good SET name=$1,price=$2,id_goodCategory=$3 where id=$4 returning id,name,price,id_goodCategory", name, prise, CategoryID, id).Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Good{}, ErrNotFound
	}
	if err != nil {
		return model.Good{}, fmt.Errorf("update good %d: %w", id, err)
	}
	return g, nil
}
func (s *PgStore) GetGood(ctx context.Context, id int) (model.Good, error) {
	var g model.Good
	err := s.db.QueryRowContext(ctx, "SELECT id, name, price, id_goodCategory FROM good WHERE id=$1", id).Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Good{}, ErrNotFound
	}
	if err != nil {
		return model.Good{}, fmt.Errorf("get good %d: %w", id, err)
	}
	return g, nil
}

func (s *PgStore) CreateGood(ctx context.Context, name string, prise float64, CategoryID *int) (model.Good, error) {
	var g model.Good
	err := s.db.QueryRowContext(ctx, "INSERT into good(name,price,id_goodCategory) VALUES ($1, $2, $3) RETURNING id", name, prise, CategoryID).Scan(&g.ID)
	if err != nil {
		return model.Good{}, fmt.Errorf("create good: %w", err)
	}
	g.Name = name
	g.Price = prise
	g.CategoryID = CategoryID
	return g, nil
}
func (s *PgStore) ListGoods(ctx context.Context) ([]model.Good, error) {
	good := []model.Good{}
	rows, err := s.db.QueryContext(ctx, "SELECT id,name,price,id_goodCategory from good order by id")
	if err != nil {
		return nil, fmt.Errorf("query goods:%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g model.Good
		if err := rows.Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID); err != nil {
			return nil, err
		}
		good = append(good, g)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return good, nil
}

// Проверка при компиляции
var _ GoodStore = (*PgStore)(nil)
