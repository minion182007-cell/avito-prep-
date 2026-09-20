package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	pingHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusOK)
	}

	want := `{"status":"ok"}`
	got := strings.TrimSpace(rec.Body.String())
	if got != want {
		t.Errorf("got body %q, want %q", got, want)
	}
}
