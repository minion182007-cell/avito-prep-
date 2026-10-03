package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHeadler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	pingHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusOK)
	}
	got := strings.TrimSpace(rec.Body.String())
	want := `{"status":"ok"}`
	if got != want {
		t.Errorf("got body %q, want %q", got, want)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
}
