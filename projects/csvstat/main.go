package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/minion182007-cell/avito-prep/projects/csvstat/internal/stats"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: csvstat <file> <column>")
		os.Exit(1)
	}
	path := os.Args[1]
	column := os.Args[2]
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "error: file not found:", path)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	defer f.Close()
	values, err := stats.ReadColumn(f, column)
	if err != nil {
		switch {
		case errors.Is(err, stats.ErrSearchColumn):
			fmt.Fprintln(os.Stderr, "error: no such column:", err)
		case errors.Is(err, stats.ErrIsEmpty):
			fmt.Fprintln(os.Stderr, "error: file is empty:", path)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
	stats := stats.ComputeStats(values)
	fmt.Println("stats:", stats.Rows)
	fmt.Println("unique:", stats.Unique)
	if stats.Numeric {
		fmt.Println("sum:", stats.Sum)
		fmt.Println("avg:", stats.Avg)
		fmt.Println("max:", stats.Max)
		fmt.Println("min:", stats.Min)
	}
}
