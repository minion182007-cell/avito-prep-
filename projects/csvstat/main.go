package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Stats struct {
	Rows    int
	Unique  int
	Numeric bool
	Sum     float64
	Avg     float64
	Min     float64
	Max     float64
}

func computeStats(values []string) Stats {
	var stats Stats
	v := len(values)
	var t []string
	for _, v := range values {
		if !slices.Contains(t, v) && v != "" {
			t = append(t, v)
		}
	}
	stats.Rows = v
	stats.Unique = len(t)
	stats.Numeric = false
	sum := 0.0
	var min, max float64
	for i, v := range values {
		v64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return stats
		}
		if i == 0 {
			max = v64
			min = v64
		} else {
			if v64 < min {
				min = v64
			}
			if v64 > max {
				max = v64
			}
		}
		sum += v64
	}
	stats.Numeric = true
	stats.Sum = sum
	stats.Avg = sum / float64(len(t))
	stats.Max = max
	stats.Min = min
	return stats

}
func readColumn(path, column string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s:%w", path, err)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv:%w", err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	idx := -1
	for i, v := range records[0] {
		if column == v {
			idx = i
			break
		}

	}
	if idx == -1 {
		return nil, fmt.Errorf("column %q not found, available: %v", column, records[0])
	}
	var values []string
	for _, row := range records[1:] {
		values = append(values, row[idx])
	}
	return values, nil
}
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: csvstat <file> <column>")
		os.Exit(1)
	}
	path := os.Args[1]
	column := os.Args[2]
	values, err := readColumn(path, column)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	stats := computeStats(values)
	fmt.Println("rows:", stats.Rows)
	fmt.Println("unique:", stats.Unique)

	if stats.Numeric {
		fmt.Println("sum:", stats.Sum)
		fmt.Println("avg:", stats.Avg)
		fmt.Println("max:", stats.Max)
		fmt.Println("min:", stats.Min)
	}

}
