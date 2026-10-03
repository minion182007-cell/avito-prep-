package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type PingResponse struct {
	Status string `json:"status"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Status: "ok"})
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("listening %q", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}
