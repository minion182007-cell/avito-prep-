package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type strucrResponse struct {
	Struct string `status:json`
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func pingHendler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, strucrResponse{Struct: "ok"})
}
func healthHeadler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, strucrResponse{Struct: "healthy"})
}
func Newmux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET/ ping", pingHendler)
	mux.HandleFunc("GET/ healthy", healthHeadler)
	return mux
}
func main() {
	post := os.Getenv("POST")
	if post == "" {
		post = "8080"
	}
	add := ":" + post
	log.Printf("addr is %s", add)
	log.Fatal(http.ListenAndServe(add, Newmux()))
}
