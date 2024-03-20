package main

import (
	"flag"
	"net/http"

	"github.com/HsiaoCz/code-beast/sth/gofolder/api"
	"github.com/gorilla/mux"
)

func main() {
	listenAddr := flag.String("listenAddr", ":50001", "listen address")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", api.HandleGetUser).Methods("GET")
	r.HandleFunc("/account/{id}", api.HandleGetAccount).Methods("GET")
	http.ListenAndServe(*listenAddr, r)
}
