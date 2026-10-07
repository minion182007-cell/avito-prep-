package stats

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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

func ReadColumn(r io.Reader, column string) ([]string, error) {
	recorder, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
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

func ComputeStats(values []string) Stats {
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
