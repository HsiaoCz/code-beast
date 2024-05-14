package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var (
		logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{}))
		port   = os.Getenv("PORT")
		router = http.NewServeMux()
		srv    = http.Server{
			Handler:      router,
			Addr:         port,
			ReadTimeout:  time.Millisecond * 1500,
			WriteTimeout: time.Millisecond * 1500,
		}
	)
	slog.SetDefault(logger)

	router.HandleFunc("/data", HandleData)

	slog.Info("the http server is running on", "port", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve http server failed %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("http server shutdown failed %v\n", err)
	}
}
func HandleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", data)
}
