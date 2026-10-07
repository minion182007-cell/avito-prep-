package stats

import (
	"encoding/csv"
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want Stats
	}{{"numbers", []string{"23", "31", "27"}, Stats{Rows: 3, Unique: 3, Numeric: true, Sum: 81, Avg: 27, Min: 23, Max: 31}},
		{"text", []string{"Moscow", "Kazan", "Moscow"}, Stats{Rows: 3, Unique: 2}},
		{"empty", []string{}, Stats{}},
		{"negative", []string{"-5", "10", "-2"}, Stats{Rows: 3, Unique: 3, Numeric: true, Sum: 3, Avg: 1, Min: -5, Max: 10}},
		{"duplicates", []string{"7", "7", "7"}, Stats{Rows: 3, Unique: 1, Numeric: true, Sum: 21, Avg: 7, Min: 7, Max: 7}},
		{"unsorted", []string{"3", "1", "2"}, Stats{Rows: 3, Sum: 6, Avg: 2, Max: 3, Min: 1, Unique: 3, Numeric: true}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeStats(tt.in)
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
func TestError(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		column  string
		wantErr error
	}{{"empty file", "", "age", ErrIsEmpty},
		{"no column", "testdata/data.csv", "agee", ErrSearchColumn},
		{"broken csv", "a,b\n1\n", "a", csv.ErrFieldCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadColumn(strings.NewReader(tt.input), tt.column)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
func TestReadColumnOK(t *testing.T) {
	r := strings.NewReader("name,age\nAnn,23\nBoris,31\n")
	got, err := ReadColumn(r, "age")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"23", "31"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
