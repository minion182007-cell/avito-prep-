package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type pingResponse struct {
	Status string `json:"status"`
}

func printJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func pingHendler(w http.ResponseWriter, r *http.Request) {
	printJSON(w, pingResponse{Status: "ok"})
}
func pingHealth(w http.ResponseWriter, r *http.Request) {
	printJSON(w, pingResponse{Status: "healthy"})
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHendler)
	mux.HandleFunc("GET /health", pingHealth)
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	add := ":" + port
	mux := Mux()
	log.Fatal(http.ListenAndServe(add, mux))
}
