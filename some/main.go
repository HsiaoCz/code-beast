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
	router := http.NewServeMux()
	userHandler := handlers.NewUserHandler(store)
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
