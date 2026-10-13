package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHeadler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantBody string
	}{
		{"ping ok", http.MethodGet, "/ping", http.StatusOK, `"status":"ok"`},
		{"health ok", http.MethodGet, "/health", http.StatusOK, `"status":"healthy"`},
		{"ping wrong method", http.MethodPatch, "/ping", http.StatusMethodNotAllowed, ""},
		{"ping ok", http.MethodGet, "/ping", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			res := httptest.NewRecorder()

			Newmux().ServeHTTP(res, req)
			if res.Code != tt.wantCode {
				t.Fatal("got status %d want %d", res.Code, tt.wantCode)
			}
			if tt.wantBody == "" {
				return
			}
			if got := strings.TrimSpace(res.Body.String()); got != tt.wantBody {
				t.Errorf("got body %q, want %q", got, tt.wantBody)
			}
			if ct := res.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q, want %q", ct, "application/json")
			}
		})
	}
}
func TestWrongMethodAllowHeader(t *testing.T)  {
	req:=httptest.NewRequest(http.MethodGet,"/ping",nil)
	rec:=httptest.NewRecorder()
	Newmux().ServeHTTP(rec,req)
	if allow:=rec.Header().Get("Allow");allow!="GET, HEAD"{
		t.Errorf("Allow = %q,want%q",allow,"GET, HEAD")
	}

}