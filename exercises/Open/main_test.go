package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestParseAge(t *testing.T) {

	n, err := parseAge("23")
	if err != nil {
		t.Errorf("parseAge(\"23\"): unexpected error: %v", err)
	}
	if n != 23 {
		t.Errorf("parseAge(\"23\") = %d, want 23", n)
	}
	_, err2 := parseAge("")
	if !errors.Is(err2, ErrEmpty) {
		t.Errorf("%v", err2)
	}
	_, err3 := parseAge("200")
	if !errors.Is(err3, ErrOutOfRange) {
		t.Errorf("%v", err3)
	}
	_, err4 := parseAge("abc")
	if !errors.Is(err4, strconv.ErrSyntax) {
		t.Errorf("%v", err4)
	}

}
