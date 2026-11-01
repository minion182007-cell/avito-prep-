package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/minion182007-cell/avito-prep/exercises/repeat-week7/csvstat/internal/stats"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "csv usage: <file> <column>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "file not found")
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	defer f.Close()
	values, err := stats.ReadColumn(f, os.Args[2])
	if err != nil {
		var ce *stats.ColumnError
		switch {
		case errors.As(err, &ce):
			fmt.Fprintf(os.Stderr, "erorr colume: %q and avilable: %s\n", ce.Column, strings.Join(ce.Avialable, ", "))
		case errors.Is(err, stats.ErrFileIsEmpty):
			fmt.Fprintln(os.Stderr, "File is empty")
		default:
			fmt.Fprintln(os.Stderr, err)

		}
		os.Exit(1)
	}
	s := stats.Compute(values)
	fmt.Println(s.Rows)
	fmt.Println(s.Uniq)
	if s.Numeric {
		fmt.Println("sum:", s.Sum)
		fmt.Println("avg:", s.Avg)
		fmt.Println("max:", s.Max)
		fmt.Println("min:", s.Min)
	}

}
