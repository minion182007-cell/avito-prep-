package server

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/model"
	"github.com/minion182007-cell/avito-prep/projects/goodsapi/internal/store"
)

func TestHandlersFake(t *testing.T) {
	tests := []struct {
		name     string
		store    store.GoodStore
		method   string
		path     string
		body     string
		code     int
		wantBody string
	}{
		// --- подделка с данными ---
		{"get ok", store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}),
			http.MethodGet, "/goods/1", "", http.StatusOK, `"name":"Phone"`},
		{"get not found", store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}),
			http.MethodGet, "/goods/99", "", http.StatusNotFound, ""},
		{"list", store.NewFakeStore(model.Good{ID: 1, Name: "Phone", Price: 100}, model.Good{ID: 2, Name: "Mouse", Price: 50}),
			http.MethodGet, "/goods", "", http.StatusOK, `"name":"Mouse"`},
		{"list empty", store.NewFakeStore(),
			http.MethodGet, "/goods", "", http.StatusOK, `[]`},
		{"create", store.NewFakeStore(),
			http.MethodPost, "/goods", `{"name":"Mouse","price":50}`, http.StatusCreated, `"id":1`},

		// --- сломанное хранилище → 500 ---
		{"get 500", &store.ErrStore{},
			http.MethodGet, "/goods/1", "", http.StatusInternalServerError, ""},
		{"list 500", &store.ErrStore{},
			http.MethodGet, "/goods", "", http.StatusInternalServerError, ""},
		{"create 500", &store.ErrStore{},
			http.MethodPost, "/goods", `{"name":"Mouse","price":50}`, http.StatusInternalServerError, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := newMux(&server{store: tt.store})
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			mux.ServeHTTP(rec, req)
			if tt.code != rec.Code {
				t.Fatalf("want code: %d got: %d", tt.code, rec.Code)
			}
			if !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf("body %q,want %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
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
		{"delete bad id", http.MethodDelete, "/goods/abc", "", http.StatusBadRequest},
		{"put bad id", http.MethodPut, "/goods/abc", `{"name":"x","price":1}`, http.StatusBadRequest},
		{"put bad json", http.MethodPut, "/goods/1", `не json`, http.StatusBadRequest},
		{"put empty name", http.MethodPut, "/goods/1", `{"name":"","price":100}`, http.StatusBadRequest},
		{"put bad price", http.MethodPut, "/goods/1", `{"name":"x","price":0}`, http.StatusBadRequest},
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
	mux := newMux(&server{store: store.NewPgStore(db)})
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
