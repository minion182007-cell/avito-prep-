package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var ErrNotFound = errors.New("good not found")

type Good struct {
	ID         int
	Name       string
	Price      float64
	CategoryID *int
}

func getGood(ctx context.Context, db *sql.DB, id int) (Good, error) {
	var g Good
	err := db.QueryRowContext(ctx, "SELECT id, name, price, id_goodCategory FROM good WHERE id=$1", id).Scan(&g.ID, &g.Name, &g.Price, &g.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return Good{}, ErrNotFound
	}
	if err != nil {
		return Good{}, fmt.Errorf("get good %d: %w", id, err)
	}
	return g, err
}
func ListenGood(ctx context.Context, db *sql.DB) ([]Good, error) {
	var good []Good
	rows, err := db.QueryContext(ctx, "SELECT id,name,price,id_goodCategory from good order by id")
	if err != nil {
		return nil, fmt.Errorf("error is:%w", err)
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
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: goods <id>")
		os.Exit(1)
	}
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
	//блок практики 2
	good, err := ListenGood(ctx, db)
	if err != nil {
		log.Fatal("error Good:", err)
	}
	for _, i := range good {
		if i.CategoryID == nil {
			fmt.Println(i.ID, i.Name, i.Price, "без категории")
		} else {
			fmt.Println(i.ID, i.Name, i.Price, *i.CategoryID)

		}

	}

	//блок практики 3

	s, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("id must be a number")
	}
	i, err := getGood(ctx, db, s)
	if errors.Is(err, ErrNotFound) {
		fmt.Fprintf(os.Stderr, "товар %d не найден\n", s)
		os.Exit(1)
	}
	if err != nil {
		log.Fatal("get good: ", err)
	}

	if i.CategoryID == nil {
		fmt.Println(i.ID, i.Name, i.Price, "без категории")
	} else {
		fmt.Println(i.ID, i.Name, i.Price, *i.CategoryID)

	}

}
