package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
	rt "app/handlers"
	st"app/structs"
)

func main() {
	err := gonfig.GetConf("configuration/conf.json", &st.Config)
	if err != nil {
		fmt.Println("error:", err)
		}
	r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", rt.EncryptHandler)
	r.HandleFunc("/api/decrypt", rt.DecryptHandler)
	http.Handle("/", r)
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", nil)
}
