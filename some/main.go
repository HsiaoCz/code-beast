package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/some/handlers"
	"github.com/HsiaoCz/code-beast/some/utils"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9001", "set server listen address")
	flag.Parse()
	router := http.NewServeMux()
	userHandler := handlers.NewUserHandler()
	router.HandleFunc("POST /user/create", utils.TransferHandler(userHandler.HandleCreateUser))
	router.HandleFunc("GET /user/show", utils.TransferHandler(userHandler.HandleShowUser))
	server := &http.Server{
		Addr:         *listenAddr,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
		Handler:      router,
	}
	slog.Info("the server is running", "port", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
