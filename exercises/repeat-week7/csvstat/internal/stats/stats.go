package stats

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrFileIsEmpty    = errors.New("File is empty")
	ErrColumnNotFound = errors.New("column not found")
)

type ColumnError struct {
	Column    string
	Avialable []string
}
type Stats struct {
	Rows               int
	Uniq               int
	Numeric            bool
	Max, Min, Sum, Avg float64
}

func (e *ColumnError) Error() string {
	return fmt.Sprintf("column %q avilable %q", e.Column, e.Avialable)
}
func (e *ColumnError) Unwrap() error {
	return ErrColumnNotFound
}

func ReadColumn(v io.Reader, target string) ([]string, error) {
	f, err := csv.NewReader(v).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read:%w", err)
	}
	if len(f) == 0 {
		return nil, ErrFileIsEmpty
	}
	idx := -1
	for i, v := range f[0] {
		if v == target {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, &ColumnError{Column: target, Avialable: f[0]}
	}
	answer := make([]string, 0, len(f))
	for _, v := range f[1:] {
		answer = append(answer, v[idx])
	}
	return answer, nil
}
func Compute(str []string) Stats {
	var s Stats
	if len(str) == 0 {
		return s
	}
	s.Rows = len(str)
	uni := make(map[string]bool)
	for _, v := range str {
		uni[v] = true
	}
	s.Uniq = len(uni)
	for i, v := range str {
		v64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return s
		}
		s.Sum += v64

		if v64 > s.Max || i == 0 {
			s.Max = v64
		}
		if v64 < s.Min || i == 0 {
			s.Min = v64
		}
	}
	s.Numeric = true
	s.Avg = s.Sum / float64(len(str))
	return s
}
