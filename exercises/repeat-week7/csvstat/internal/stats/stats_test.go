package stats

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
	"testing"
)

func TestReadColumn(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		column  string
		want    []string
		wanterr error
	}{
		{"normal", "name,price\nphone,100\ntable,300\n", "price", []string{"100", "300"}, nil},
		{"first column", "name,price\nphone,100\ntable,300\n", "name", []string{"phone", "table"}, nil},
		{"header only", "name,price\n", "price", []string{""}, nil},
		{"empty file", "", "price", nil, ErrFileIsEmpty},
		{"no column", "name,price\nphone,100\ntable,300\n", "qty", nil, ErrColumnNotFound},
		{"bad csv", "name,price\nphone\n", "price", nil, csv.ErrFieldCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadColumn(strings.NewReader(tt.in), tt.column)
			if !errors.Is(err, tt.wanterr) {
				log.Fatalf("error is %v", tt.wanterr)
			}
			if tt.wanterr != nil {
				return
			}
			if !slices.Equal(got, tt.want) {
				fmt.Errorf("want:%+v got:%+v", tt.want, got)
			}
		})
	}

}
func TestErrColumn_Error(t *testing.T) {
	_, err := ReadColumn(strings.NewReader("name,price\nphone,100\ntable,300\n"), "qty")
	var ce *ColumnError
	if !errors.As(err, &ce) {
		log.Fatalf("error column:%q avilable is:%s", ce.Column, ce.Avialable)

	}
	if !errors.Is(err, ErrColumnNotFound) {
		t.Fatalf("got:%v want:%v", err, ErrColumnNotFound)
	}
	if ce.Column != "qty" {
		t.Fatalf("got:%q want:%q", ce.Column, "qty")
	}
	if !slices.Equal(ce.Avialable, []string{"name", "price"}) {
		t.Fatalf("got:%+v want:%+v", ce.Avialable, []string{"name", "price"})
	}
}
func TestCompute(t *testing.T) {
	test := []struct {
		name string
		in   []string
		want Stats
	}{
		{"number", []string{"1", "2", "3"}, Stats{Rows: 3, Uniq: 3, Numeric: true, Sum: 6, Avg: 2, Min: 1, Max: 3}},
		{"non sort", []string{"5", "1", "3"}, Stats{Rows: 3, Uniq: 3, Numeric: true, Sum: 9, Avg: 3, Min: 1, Max: 5}},
		{"repeat", []string{"a", "a", "b"}, Stats{Rows: 3, Uniq: 2}},
		{"non number", []string{"a", "b"}, Stats{Rows: 2, Uniq: 2, Numeric: false}},
		{"empty", []string{}, Stats{}},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Compute(tt.in)
			if got != tt.want {
				t.Errorf("got: %+v want: %+v", got, tt.want)

			}
		})
	}
}
