package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/minion182007-cell/avito-prep/exercises/csv_prac/internal/stats"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: csvtat <file> <column>")
		os.Exit(1)
	}
	path := os.Args[1]
	column := os.Args[2]
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "file not find:", path)
		} else {
			fmt.Fprintln(os.Stderr, "err", err)
		}
		os.Exit(1)
	}

	values, err := stats.ReadColumn(f, column)
	if err != nil {
		var ce *stats.ColumnError
		switch {
		case errors.As(err, &ce):
			fmt.Fprintf(os.Stderr, "error: no column %q, available: %s\n", ce.Column, strings.Join(ce.Available, ", "))
		case errors.Is(err, stats.ErrFileIsEmpty):
			fmt.Fprintln(os.Stderr, "error: file is empty:", path)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
	s := stats.Compute(values)
	fmt.Println("rows:", s.Rows)
	fmt.Println("unique:", s.Uniq)
	if s.Numeric {
		fmt.Println("sum:", s.Sum)
		fmt.Println("avg:", s.Avg)
		fmt.Println("max:", s.Max)
		fmt.Println("min:", s.Min)

	}
}
