package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", handleGetUserByID)
	server := http.Server{
		Handler:      router,
		Addr:         ":9001",
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Microsecond * 1500,
	}
	log.Fatal(server.ListenAndServe())
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// id:=r.PathValue("id")
	fmt.Println(id)
}
