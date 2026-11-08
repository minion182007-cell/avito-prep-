package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

type server struct {
	store store.GoodStore
}

func (s *server) handleDelete(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest) //400
		return
	}
	err = s.store.DeleteGood(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "good not found", http.StatusNotFound) // 404
		return
	}
	if err != nil {
		log.Println("delete good:", err)
		http.Error(w, "internal error", http.StatusInternalServerError) // 500
		return
	}
	w.WriteHeader(204)
}
func (s *server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest) //400
		return
	}
	var in model.CreateGoodInput
	err = json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if in.Name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}
	if in.Price <= 0 {
		http.Error(w, "price must be positive", http.StatusBadRequest)
		return
	}

	goods, err := s.store.UpdateGoods(r.Context(), id, in.Name, in.Price, in.CategoryID)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "good not found", http.StatusNotFound) // 404
		return
	}
	if err != nil {
		log.Println("update good:", err)
		http.Error(w, "internal error", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}
func (s *server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var in model.CreateGoodInput
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if in.Name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}
	if in.Price <= 0 {
		http.Error(w, "price must be positive", http.StatusBadRequest)
		return
	}
	goods, err := s.store.CreateGood(r.Context(), in.Name, in.Price, in.CategoryID)
	if err != nil {
		log.Println("crate good:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(goods)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func (s *server) handleList(w http.ResponseWriter, r *http.Request) {
	goods, err := s.store.ListGoods(r.Context())
	if err != nil {
		log.Println("goods list:", err)
		http.Error(w, "internal error", http.StatusInternalServerError) //500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest) //400
		return
	}

	g, err := s.store.GetGood(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "good not found", http.StatusNotFound) // 404
		return
	}
	if err != nil {
		log.Println("get good:", err)
		http.Error(w, "internal error", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}
func NewHandler(st store.GoodStore) http.Handler {
	return newMux(&server{store: st})
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
