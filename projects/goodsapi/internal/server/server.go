package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/service"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

type GoodService interface {
	Create(ctx context.Context, in model.CreateGoodInput) (model.Good, error)
	Update(ctx context.Context, id int, in model.CreateGoodInput) (model.Good, error)
	Get(ctx context.Context, id int) (model.Good, error)
	List(ctx context.Context) ([]model.Good, error)
	Delete(ctx context.Context, id int) error
}

type server struct {
	svc GoodService
}

// writeError превращает ошибку сервиса в HTTP-статус.
func writeError(w http.ResponseWriter, op string, err error) {
	var ve *service.ValidationError
	switch {
	case errors.As(err, &ve):
		http.Error(w, ve.Msg, http.StatusBadRequest) // 400
	case errors.Is(err, service.ErrNotFound):
		http.Error(w, "good not found", http.StatusNotFound) // 404
	default:
		log.Println(op+":", err)
		http.Error(w, "internal error", http.StatusInternalServerError) // 500
	}
}

func (s *server) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}
	if err := s.svc.Delete(r.Context(), id); err != nil {
		writeError(w, "delete good", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}
	var in model.CreateGoodInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	good, err := s.svc.Update(r.Context(), id, in)
	if err != nil {
		writeError(w, "update good", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(good)
}

func (s *server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var in model.CreateGoodInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	good, err := s.svc.Create(r.Context(), in)
	if err != nil {
		writeError(w, "create good", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(good)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (s *server) handleList(w http.ResponseWriter, r *http.Request) {
	goods, err := s.svc.List(r.Context())
	if err != nil {
		writeError(w, "list goods", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}
	good, err := s.svc.Get(r.Context(), id)
	if err != nil {
		writeError(w, "get good", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(good)
}

// NewHandler собирает цепочку: store -> service -> server.
func NewHandler(st store.GoodStore) http.Handler {
	return newMux(&server{svc: service.New(st)})
}

func newMux(srv *server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", srv.handleHealth)
	mux.HandleFunc("GET /goods", srv.handleList)
	mux.HandleFunc("GET /goods/{id}", srv.handleGet)
	mux.HandleFunc("POST /goods", srv.handleCreate)
	mux.HandleFunc("DELETE /goods/{id}", srv.handleDelete)
	mux.HandleFunc("PUT /goods/{id}", srv.handleUpdate)
	return mux
}
