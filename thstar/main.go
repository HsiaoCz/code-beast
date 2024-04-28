package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/HsiaoCz/code-beast/thstar/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/foo", handlers.TransferHandlerFunc(handlers.HandleFoo))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":9001"
	}
	slog.Info("the server started", "port", port)
	srv := http.Server{
		Handler:      router,
		Addr:         port,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("the server running error", "err", err)
		return
	}
}
