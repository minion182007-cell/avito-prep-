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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err = db.PingContext(ctx);
	cancel()
	if  err != nil {
		log.Fatal("cannot connect: ", err)

	}
	handler := server.NewHandler(store.NewPgStore(db))
	log.Println("listening on :"+port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
