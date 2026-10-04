package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Stats struct {
	Rows    int
	Unique  int
	Numeric bool
	Sum     float64
	Max     float64
	Min     float64
	Avg     float64
}

func readColumn(path, column string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read %s column %w", path, err)
	}
	defer f.Close()
	recorders, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if len(recorders) == 0 {
		return nil, fmt.Errorf("File is empty")
	}
	idx := -1
	for i, v := range recorders[0] {
		if column == v {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, fmt.Errorf("%q not found,aviable %v", column, recorders[0])
	}
	var values []string
	for _, row := range recorders[1:] {
		values = append(values, row[idx])
	}
	return values, nil
}

func computeStats(values []string) Stats {
	var stats Stats
	if len(values) == 0 {
		return stats
	}

	stats.Rows = len(values)
	uniq := make(map[string]bool)
	for _, v := range values {
		if v != "" {
			uniq[v] = true
		}
	}
	stats.Unique = len(uniq)

	sum := 0.0
	var max, min float64
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
	stats.Numeric = true
	stats.Min = min
	stats.Max = max
	stats.Sum = sum
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
