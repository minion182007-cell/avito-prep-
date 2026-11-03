package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	mux := newMux(&server{})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("code = %d,want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); got != "ok" {
		t.Errorf("body %q,want %q", got, "ok")
	}
}
func TestRoutes(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		body   string
		code   int
	}{
		{"health", http.MethodGet, "/health", "", http.StatusOK},
		{"bad id", http.MethodGet, "/goods/abc", "", http.StatusBadRequest},
		{"wrong method", http.MethodPost, "/goods/1", "", http.StatusMethodNotAllowed},
		{"unknown path", http.MethodGet, "/nope", "", http.StatusNotFound},
		{"bad json", http.MethodPost, "/goods", `не json`, http.StatusBadRequest},
		{"empty name", http.MethodPost, "/goods", `{"name":"","price":100}`, http.StatusBadRequest},
		{"bad price", http.MethodPost, "/goods", `{"name":"x","price":0}`, http.StatusBadRequest},
	}
	mux := newMux(&server{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			if rec.Code != tt.code {
				t.Errorf("%s %s: code = %d, want %d", tt.method, tt.path, rec.Code, tt.code)
			}
		})
	}
}

func TestGoodsDB(t *testing.T) {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatalf("open is : %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Fatalf("ping: %v", err)
	}

	tests := []struct {
		name     string
		path     string
		wantCode int
		wantBody string
	}{
		{"existing", "/goods/1", http.StatusOK, `"name":"iPhone"`},
		{"not found", "/goods/99", http.StatusNotFound, "good not found"},
		{"list", "/goods", http.StatusOK, `"category_id":null`},
	}
	mux := newMux(&server{store: &pgStore{db: db}})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("code = %d, want %d", rec.Code, tt.wantCode)
			}
			if !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf("body = %q, want it to contain %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
