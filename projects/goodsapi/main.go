package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var ErrNotFound = errors.New("good not found")

type Good struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID *int    `json:"category_id"`
}
type createGoodInput struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID *int    `json:"category_id"`
}

// getGood возвращает ErrNotFound
type GoodStore interface {
	deleteGood(ctx context.Context, id int) error
	updateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (Good, error)
	getGood(ctx context.Context, id int) (Good, error)
	createGood(ctx context.Context, name string, prise float64, CategoryID *int) (Good, error)
	ListGoods(ctx context.Context) ([]Good, error)
}
type pgStore struct {
	db *sql.DB
}
type server struct {
	store GoodStore
}

func (s *pgStore) deleteGood(ctx context.Context, id int) error {
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
func (s *pgStore) updateGoods(ctx context.Context, id int, name string, prise float64, CategoryID *int) (Good, error) {
	var g Good
	err := s.db.QueryRowContext(ctx, "UPDATE good SET name=$1,price=$2,id_goodCategory=$3 where id=$4 returning id,name,price,id_goodCategory", name, prise, CategoryID, id).Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return Good{}, ErrNotFound
	}
	if err != nil {
		return Good{}, fmt.Errorf("update good %d: %w", id, err)
	}
	return g, nil
}
func (s *pgStore) getGood(ctx context.Context, id int) (Good, error) {
	var g Good
	err := s.db.QueryRowContext(ctx, "SELECT id, name, price, id_goodCategory FROM good WHERE id=$1", id).Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return Good{}, ErrNotFound
	}
	if err != nil {
		return Good{}, fmt.Errorf("get good %d: %w", id, err)
	}
	return g, nil
}

func (s *pgStore) createGood(ctx context.Context, name string, prise float64, CategoryID *int) (Good, error) {
	var g Good
	err := s.db.QueryRowContext(ctx, "INSERT into good(name,price,id_goodCategory) VALUES ($1, $2, $3) RETURNING id", name, prise, CategoryID).Scan(&g.ID)
	if err != nil {
		return Good{}, fmt.Errorf("create good: %w", err)
	}
	g.Name = name
	g.Price = prise
	g.CategoryID = CategoryID
	return g, nil
}
func (s *pgStore) ListGoods(ctx context.Context) ([]Good, error) {
	good := []Good{}
	rows, err := s.db.QueryContext(ctx, "SELECT id,name,price,id_goodCategory from good order by id")
	if err != nil {
		return nil, fmt.Errorf("query goods:%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g Good
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
var _ GoodStore = (*pgStore)(nil)

func (s *server) handleDelete(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest) //400
		return
	}
	err = s.store.deleteGood(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
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
	var in createGoodInput
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
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest) //400
		return
	}
	goods, err := s.store.updateGoods(r.Context(), id, in.Name, in.Price, in.CategoryID)
	if errors.Is(err, ErrNotFound) {
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
	var in createGoodInput
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
	goods, err := s.store.createGood(r.Context(), in.Name, in.Price, in.CategoryID)
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

	g, err := s.store.getGood(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
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

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect: ", err)
	}
	srv := &server{store: &pgStore{db: db}}
	mux := newMux(srv)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
