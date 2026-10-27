package stats

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestReadColumn(t *testing.T) {
	test := []struct {
		name     string
		input    string
		column   string
		want     []string
		want_err error
	}{
		{"normal", "name,price\nphone,100\ntable,300\n", "price", []string{"100", "300"}, nil},
		{"first column", "name,price\nphone,100\ntable,300\n", "name", []string{"phone", "table"}, nil},
		{"title", "name,price", "price", []string{}, nil},
		{"Empty file", "", "price", nil, ErrFileIsEmpty},
		{"No title", "phone,100\ntable,300\n", "qty", nil, ErrColumnNotFound},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadColumn(strings.NewReader(tt.input), tt.column)
			if !errors.Is(err, tt.want_err) {
				t.Fatalf("got: %v want: %v", err, tt.want_err)
			}
			if tt.want_err != nil {
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got: %q want: %+v", got, tt.want)
			}
		})
	}
}
func TestColumnError_Error(t *testing.T) {
	_, err := ReadColumn(strings.NewReader("name,price\nphone,100\ntable,300\n"), "qty")
	var ce *ColumnError
	if !errors.As(err, &ce) {
		t.Fatal(err)
	}
	if !errors.Is(err, ErrColumnNotFound) {
		t.Fatalf("got: %v want: %v", err, ErrColumnNotFound)
	}
	if ce.Column != "qty" {
		t.Errorf("want:%q got:%q", "qty", ce.Column)
	}
	if !slices.Equal(ce.Available, []string{"name", "price"}) {
		t.Errorf("want:%+v got:%+v", []string{"name", "price"}, ce.Available)
	}

}
func TestCompute(t *testing.T) {
	test := []struct {
		name string
		in   []string
		want Stats
	}{
		{"number", []string{"1", "2", "3"}, Stats{Rows: 3, Uniq: 3, Numeric: true, Sum: 6, Avg: 2, Min: 1, Max: 3}},
		{"non sort", []string{"5", "1", "3"}, Stats{Rows: 3, Uniq: 3, Numeric: true, Sum: 6, Avg: 3, Min: 1, Max: 5}},
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
