package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib"

)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL wrong")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect:", err)
	}

	var count int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM good").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("товаров в базе:", count)
}
