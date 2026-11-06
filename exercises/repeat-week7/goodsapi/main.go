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
)

var Errnotfound = errors.New("good not found")

type Good struct {
	Name        string  `json:"name"`
	Id          int     `json:"id"`
	Price       float64 `json:"price"`
	Category_ID *int    `json:"category_id"`
}

type CategoryGood struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category_ID *int    `json:"category_id"`
}
type server struct {
	db *sql.DB
}

func GetGood(ctx context.Context, db *sql.DB, id int) (Good, error) {
	var g Good
	err := db.QueryRowContext(ctx, "SELECT id,name,price,id_goodCategory from good where id=$1", id).Scan(&g.Id, &g.Name, &g.Price, &g.Category_ID)
	if errors.Is(err, sql.ErrNoRows) {
		return Good{}, Errnotfound
	}
	if err != nil {
		return Good{}, fmt.Errorf("get good:%d err:%w", id, err)
	}
	return g, nil
}

func CreateGood(ctx context.Context, db *sql.DB, name string, price float64, category_id *int) (Good, error) {
	var g Good
	err := db.QueryRowContext(ctx, "INSERT into good(name,price,id_goodCategory) values ($1, $2, $3) returning id", name, price, category_id).Scan(&g.Id)
	if err != nil {
		return Good{}, fmt.Errorf("create goods:%w", err)
	}
	g.Name = name
	g.Price = price
	g.Category_ID = category_id
	return g, nil
}
func ListGoods(ctx context.Context, db *sql.DB) ([]Good, error) {
	goods := []Good{}
	rows, err := db.QueryContext(ctx, "SELECT id,name,price,id_goodCategory from good order by id")
	if err != nil {
		return nil, fmt.Errorf("good query:%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g Good
		if err := rows.Scan(&g.Id, &g.Name, &g.Price, &g.Category_ID); err != nil {
			return nil, err
		}
		goods = append(goods, g)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return goods, nil
}

func (s *server) heandlCreate(w http.ResponseWriter, r *http.Request) {
	var in CategoryGood
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
		http.Error(w, "price must be pozitive", http.StatusBadRequest)
		return
	}

	goods, err := CreateGood(r.Context(), s.db, in.Name, in.Price, in.Category_ID)
	if err != nil {
		log.Println("create good", err)
		http.Error(w, "error create", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)

}

func (s *server) heandlHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (s *server) heandleList(w http.ResponseWriter, r *http.Request) {
	goods, err := ListGoods(r.Context(), s.db)
	if err != nil {
		log.Println("list goods", err)
		http.Error(w, "erorr get list", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goods)
}

func (s *server) heandlerGet(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "is not number", http.StatusBadRequest)
		return
	}
	g, err := GetGood(r.Context(), s.db, id)
	if errors.Is(err, Errnotfound) {
		log.Println("try get", err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("get goods", err)
		http.Error(w, "get good", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func Mux(srv *server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", srv.heandlHealth)
	mux.HandleFunc("GET /goods", srv.heandleList)
	mux.HandleFunc("GET /goods/{id}", srv.heandlerGet)
	mux.HandleFunc("POST /goods", srv.heandlCreate)
	return mux
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("database is empty")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("cannot connect: ", err)
	}
	defer db.Close()

	ctx, cansel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cansel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect: ", err)
	}

	srv := &server{db: db}
	mux := Mux(srv)
	log.Println("listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
