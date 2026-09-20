package main

import (
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func main() {
	log.Println("start")
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
