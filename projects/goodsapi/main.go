package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/server"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

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
	mux := server.NewHandler(store.NewPgStore(db))
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
