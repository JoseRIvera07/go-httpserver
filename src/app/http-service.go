package main

import (
	"net/http"
	"github.com/gorilla/mux"
	rt "app/routing"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", rt.Encrypt_Routing).Methods("POST")
	r.HandleFunc("/api/decrypt", rt.Decrypt_Routing).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
