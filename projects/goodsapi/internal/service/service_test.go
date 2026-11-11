package service

import (
	"context"
	"errors"
	"testing"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

// validationMsg достаёт текст из *ValidationError; ok=false, если ошибка другого вида.
func validationMsg(err error) (msg string, ok bool) {
	var ve *ValidationError
	if errors.As(err, &ve) {
		return ve.Msg, true
	}
	return "", false
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		in      model.CreateGoodInput
		wantMsg string // не пусто — ждём ValidationError с этим текстом
	}{
		{"ok", model.CreateGoodInput{Name: "Phone", Price: 100}, ""},
		{"empty name", model.CreateGoodInput{Name: "", Price: 100}, "name is empty"},
		{"zero price", model.CreateGoodInput{Name: "Phone", Price: 0}, "price must be positive"},
		{"negative price", model.CreateGoodInput{Name: "Phone", Price: -5}, "price must be positive"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := New(store.NewFakeStore())
			got, err := svc.Create(context.Background(), tt.in)

			if tt.wantMsg != "" {
				msg, ok := validationMsg(err)
				if !ok {
					t.Fatalf("want ValidationError, got %v", err)
				}
				if msg != tt.wantMsg {
					t.Errorf("msg = %q, want %q", msg, tt.wantMsg)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.ID != 1 || got.Name != tt.in.Name || got.Price != tt.in.Price {
				t.Errorf("got %+v, want id=1 name=%q price=%v", got, tt.in.Name, tt.in.Price)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		in      model.CreateGoodInput
		wantMsg string // ждём ValidationError
		wantErr error  // ждём эту ошибку через errors.Is
	}{
		{"ok", 1, model.CreateGoodInput{Name: "Phone 2", Price: 150}, "", nil},
		{"empty name", 1, model.CreateGoodInput{Name: "", Price: 150}, "name is empty", nil},
		{"bad price", 1, model.CreateGoodInput{Name: "Phone 2", Price: -1}, "price must be positive", nil},
		{"not found", 99, model.CreateGoodInput{Name: "Phone 2", Price: 150}, "", ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := New(store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}))
			got, err := svc.Update(context.Background(), tt.id, tt.in)

			switch {
			case tt.wantMsg != "":
				msg, ok := validationMsg(err)
				if !ok {
					t.Fatalf("want ValidationError, got %v", err)
				}
				if msg != tt.wantMsg {
					t.Errorf("msg = %q, want %q", msg, tt.wantMsg)
				}
			case tt.wantErr != nil:
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("err = %v, want %v", err, tt.wantErr)
				}
			default:
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got.ID != tt.id || got.Name != tt.in.Name || got.Price != tt.in.Price {
					t.Errorf("got %+v", got)
				}
			}
		})
	}
}

func TestGetNotFound(t *testing.T) {
	svc := New(store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}))
	if _, err := svc.Get(context.Background(), 99); !errors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	svc := New(store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}))

	if err := svc.Delete(ctx, 99); !errors.Is(err, ErrNotFound) {
		t.Errorf("delete missing: err = %v, want ErrNotFound", err)
	}
	if err := svc.Delete(ctx, 1); err != nil {
		t.Fatalf("delete existing: %v", err)
	}
	if _, err := svc.Get(ctx, 1); !errors.Is(err, ErrNotFound) {
		t.Errorf("after delete: err = %v, want ErrNotFound", err)
	}
}

// Валидация идёт ДО обращения к хранилищу: даже с «упавшей» базой ответ — ValidationError.
func TestValidationBeforeStore(t *testing.T) {
	svc := New(&store.ErrStore{})
	bad := model.CreateGoodInput{Name: "", Price: 100}

	if _, err := svc.Create(context.Background(), bad); !isValidation(err) {
		t.Errorf("create: err = %v, want ValidationError", err)
	}
	if _, err := svc.Update(context.Background(), 1, bad); !isValidation(err) {
		t.Errorf("update: err = %v, want ValidationError", err)
	}
}

// Ошибка хранилища проходит наверх как есть: не ValidationError и не ErrNotFound (хендлер ответит 500).
func TestStoreErrorPassesThrough(t *testing.T) {
	svc := New(&store.ErrStore{})
	_, err := svc.Create(context.Background(), model.CreateGoodInput{Name: "Phone", Price: 100})

	if err == nil {
		t.Fatal("want error, got nil")
	}
	if isValidation(err) {
		t.Error("store error must not be ValidationError")
	}
	if errors.Is(err, ErrNotFound) {
		t.Error("store error must not be ErrNotFound")
	}
}

func isValidation(err error) bool {
	_, ok := validationMsg(err)
	return ok
}
