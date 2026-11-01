package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantBody string
	}{
		{"ping", http.MethodGet, "/ping", http.StatusOK, `{"status":"ok"}`},
		{"health", http.MethodGet, "/health", http.StatusOK, `{"status":"healthy"}`},
		{"Patch", http.MethodPatch, "/ping", http.StatusMethodNotAllowed, ""},
		{"ping", http.MethodGet, "/nope", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			Mux().ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("want code: %d have: %d", tt.wantCode, rec.Code)
			}
			if tt.wantBody == "" {
				return
			}
			if bd := strings.TrimSpace(rec.Body.String()); bd != tt.wantBody {
				t.Errorf("want body: %q have: %q", tt.wantBody, bd)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("want Content-Type: %q have: %q", "application/json", ct)
			}
		})
	}
}
