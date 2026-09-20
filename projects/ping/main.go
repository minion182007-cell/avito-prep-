package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type pingResponse struct {
	Status string `json:"status"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pingResponse{Status: "ok"})
}
func main() {
	log.Println("start")
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
