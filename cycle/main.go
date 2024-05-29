package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/code-beast/cycle/app"
	"github.com/HsiaoCz/code-beast/cycle/contact"
	"github.com/HsiaoCz/code-beast/justic/handlers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("load env error", "err", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		slog.Error("connect mongo database error", "err msg", err)
		return
	}

	var (
		userColl         = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLLNAME"))
		mongoUserContact = contact.NewMongoUserContact(client, userColl)
		contact          = &contact.Dependenec{User: mongoUserContact}
		userHandler      = app.NewUserHandler(contact)
	)

	router := http.NewServeMux()

	router.HandleFunc("POST /user", handlers.TransferHandlerFunc(userHandler.HandleCreateUser))
	router.HandleFunc("GET /user", handlers.TransferHandlerFunc(userHandler.HandleGetUserByID))

	srv := http.Server{
		Handler:      router,
		Addr:         os.Getenv("PORT"),
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}

	slog.Info("the server is running", "port", os.Getenv("PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("the server shutdown error", "err", err)
		return
	}

	slog.Info("http server shutdown")
}
