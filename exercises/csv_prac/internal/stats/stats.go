package stats

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrFileIsEmpty    = errors.New("file is empty")
	ErrColumnNotFound = errors.New("column not found")
)

type ColumnError struct {
	Column    string
	Available []string
}

func (e *ColumnError) Error() string {
	return fmt.Sprintf("column: %q avilable: %q", e.Column, e.Available)
}

func (e *ColumnError) Unwrap() error { return ErrColumnNotFound }

type Stats struct {
	Rows               int
	Uniq               int
	Numeric            bool
	Max, Min, Sum, Avg float64
}

func ReadColumn(r io.Reader, column string) ([]string, error) {
	f, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv %w", err)
	}
	if len(f) == 0 {
		return nil, ErrFileIsEmpty
	}
	idx := -1
	for i, name := range f[0] {
		if name == column {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, &ColumnError{Column: column, Available: f[0]}
	}
	values := make([]string, 0, len(f)-1)
	for _, row := range f[1:] {
		values = append(values, row[idx])
	}
	return values, nil
}
func Compute(values []string) Stats {
	var s Stats
	if len(values) == 0 {
		return s
	}
	s.Rows = len(values)
	uniq := make(map[string]bool)
	for _, name := range values {
		uniq[name] = true
	}
	s.Uniq = len(uniq)
	for i, name := range values {
		x, err := strconv.ParseFloat(name, 64)
		if err != nil {
			return s
		}
		s.Sum += x
		if i == 0 || x < s.Min {
			s.Min = x
		}
		if i == 0 || x > s.Max {
			s.Max = x
		}

	}
	s.Numeric = true
	s.Avg = s.Sum / float64(s.Rows)
	return s
}
