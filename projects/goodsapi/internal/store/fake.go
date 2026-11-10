package store

import (
	"context"
	"errors"
	"sort"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
)

type fakeStore struct {
	nextID int
	mp     map[int]model.Good
}

func NewFakeStore(goods ...model.Good) *fakeStore {
	mapa := make(map[int]model.Good)
	fs := &fakeStore{nextID: 1, mp: mapa}
	for _, v := range goods {
		mapa[v.ID] = v
		if fs.nextID <= v.ID {
			fs.nextID = v.ID + 1
		}
	}
	return fs
}
func (s *fakeStore) DeleteGood(ctx context.Context, id int) error {
	_, ok := s.mp[id]
	if !ok {
		return ErrNotFound
	}
	delete(s.mp, id)
	return nil
}
func (s *fakeStore) UpdateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (model.Good, error) {
	_, ok := s.mp[id]
	if !ok {
		return model.Good{}, ErrNotFound
	}

	var g model.Good
	g.CategoryID = CategoryID
	g.Price = prise
	g.Name = name
	g.ID = id
	s.mp[id] = g
	return g, nil
}

// GetGood возвращает ErrNotFound
func (s *fakeStore) GetGood(ctx context.Context, id int) (model.Good, error) {
	g, ok := s.mp[id]
	if !ok {
		return model.Good{}, ErrNotFound
	}
	return g, nil
}

func (s *fakeStore) ListGoods(ctx context.Context) ([]model.Good, error) {
	good := []model.Good{}
	for _, g := range s.mp {
		good = append(good, g)
	}
	sort.Slice(good, func(i, j int) bool { return good[i].ID < good[j].ID })
	return good, nil

}

func (s *fakeStore) CreateGood(ctx context.Context, name string, prise float64, CategoryID *int) (model.Good, error) {
	g := model.Good{ID: s.nextID, Name: name, Price: prise, CategoryID: CategoryID}
	s.nextID++
	s.mp[g.ID] = g
	return g, nil
}
func (s *fakeStore) PurchaseGood(ctx context.Context, idShop int, idGood int, count int) error {
	_, ok := s.mp[idGood]
	if !ok {
		return ErrNotFound
	}
	return nil
}

var _ GoodStore = (*fakeStore)(nil)

type ErrStore struct{}

func (s *ErrStore) GetGood(ctx context.Context, id int) (model.Good, error) {
	return model.Good{}, errors.New("db down")
}
func (s *ErrStore) ListGoods(ctx context.Context) ([]model.Good, error) {
	return nil, errors.New("db down")
}
func (s *ErrStore) CreateGood(ctx context.Context, name string, prise float64, CategoryID *int) (model.Good, error) {
	return model.Good{}, errors.New("db down")
}
func (s *ErrStore) DeleteGood(ctx context.Context, id int) error {
	return errors.New("db down")
}
func (s *ErrStore) UpdateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (model.Good, error) {
	return model.Good{}, errors.New("db down")
}
func (s *ErrStore) PurchaseGood(ctx context.Context, idShop int, idGood int, count int) error {
	return errors.New("db down")
}

var _ GoodStore = (*ErrStore)(nil)
