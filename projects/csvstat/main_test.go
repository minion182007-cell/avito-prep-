package main

import (
	"testing"
)

func TestCompuetrNumeric(t *testing.T) {
	got := computeStats([]string{"23", "31", "27"})
	want := Stats{
		Rows:    3,
		Unique:  3,
		Numeric: true,
		Sum:     81,
		Avg:     27,
		Min:     23,
		Max:     31,
	}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestComputeStatsText(t *testing.T) {
	got := computeStats([]string{"Moscow", "Kazan", "Moscow"})
	want := Stats{
		Rows:    3,
		Unique:  2,
		Numeric: false,
	}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
