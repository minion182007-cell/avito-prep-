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

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pingResponse{Status: "ok"})
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("listening on %s", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("Get /ping", PingHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}
