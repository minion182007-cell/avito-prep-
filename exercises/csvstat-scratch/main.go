package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
)

var (
	ErrIsEmpty      = errors.New("file is empty")
	ErrSearchColumn = errors.New("Column not search")
)

type Stats struct {
	Rows    int
	Avg     float64
	Unique  int
	Sum     float64
	Max     float64
	Min     float64
	Numeric bool
}

func readColumn(r io.Reader, column string) ([]string, error) {
	recorder, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(recorder) == 0 {
		return nil, ErrIsEmpty
	}
	idx := -1
	for i, v := range recorder[0] {
		if v == column {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, fmt.Errorf("in %q: %w, avilable %q", column, ErrSearchColumn, recorder[0])
	}
	var values []string
	for _, row := range recorder[1:] {
		values = append(values, row[idx])

	}
	return values, nil

}

func computeStats(values []string) Stats {
	var stats Stats
	val := len(values)
	if val == 0 {
		return stats
	}
	stats.Rows = val
	sum := 0.0
	var max float64
	var min float64
	uniq := make(map[string]bool)
	for _, v := range values {
		if v != "" {
			uniq[v] = true
		}
	}
	stats.Unique = len(uniq)
	for i, v := range values {
		v64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return stats
		}
		sum += v64
		if i == 0 {
			max = v64
			min = v64
		} else {
			if max < v64 {
				max = v64
			}
			if min > v64 {
				min = v64
			}
		}
	}
	stats.Sum = sum

	stats.Max = max
	stats.Min = min
	stats.Numeric = true
	stats.Avg = sum / float64(len(values))
	return stats
}

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
	values, err := readColumn(f, column)
	if err != nil {
		switch {
		case errors.Is(err, ErrSearchColumn):
			fmt.Fprintln(os.Stderr, "error: no such column:", err)
		case errors.Is(err, ErrIsEmpty):
			fmt.Fprintln(os.Stderr, "error: file is empty:", path)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
	stats := computeStats(values)
	fmt.Println(stats.Rows)
	fmt.Println(stats.Unique)
	if stats.Numeric {
		fmt.Println("sum:", stats.Sum)
		fmt.Println("avg:", stats.Avg)
		fmt.Println("max:", stats.Max)
		fmt.Println("min:", stats.Min)
	}
}
