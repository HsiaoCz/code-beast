package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-beast/some/handlers"
	"github.com/HsiaoCz/code-beast/some/store"
	"github.com/HsiaoCz/code-beast/some/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURL))
	if err != nil {
		slog.Error("mongo db connection error", "err", err)
		return
	}
	listenAddr := flag.String("listenAddr", ":9001", "set server listen address")
	flag.Parse()
	var (
		userStore = store.NewMongoStore(client)
		store     = &store.Store{User: userStore}
	)
	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	userHandler := handlers.NewUserHandler(store)
	router.Post("/user/create", utils.TransferHandler(userHandler.HandleCreateUser))
	router.Get("/user/show", utils.TransferHandler(userHandler.HandleShowUser))
	server := &http.Server{
		Addr:         *listenAddr,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
		Handler:      router,
	}
	slog.Info("the server is running", "port", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
