package service

import (
	"context"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

var ErrNotFound = store.ErrNotFound

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string { return e.Msg }

type Service struct {
	store store.GoodStore
}

func New(st store.GoodStore) *Service {
	return &Service{store: st}
}
func validate(in model.CreateGoodInput) error {
	if in.Name == "" {
		return &ValidationError{Msg: "name is empty"}
	}
	if in.Price <= 0 {
		return &ValidationError{Msg: "price must be positive"}
	}
	return nil
}

func (s *Service) Create(ctx context.Context, in model.CreateGoodInput) (model.Good, error) {
	if err := validate(in); err != nil {
		return model.Good{}, err
	}
	return s.store.CreateGood(ctx, in.Name, in.Price, in.CategoryID)
}
func (s *Service) Update(ctx context.Context, id int, in model.CreateGoodInput) (model.Good, error) {
	if err := validate(in); err != nil {
		return model.Good{}, err
	}
	return s.store.UpdateGoods(ctx, id, in.Name, in.Price, in.CategoryID)
}
func (s *Service) Get(ctx context.Context, id int) (model.Good, error) {
	return s.store.GetGood(ctx, id)
}
func (s *Service) List(ctx context.Context) ([]model.Good, error) {
	return s.store.ListGoods(ctx)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.store.DeleteGood(ctx, id)
}
