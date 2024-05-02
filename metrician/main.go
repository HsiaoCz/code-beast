package main

import (
	"context"
	"fmt"
	"io"
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
		log.Fatalf("load env failed %v\n", err)
	}
	var (
		port   = os.Getenv("PORT")
		router = http.NewServeMux()
		srv    = http.Server{
			Handler:      router,
			Addr:         port,
			ReadTimeout:  time.Millisecond * 1500,
			WriteTimeout: time.Millisecond * 1500,
		}
	)

	router.HandleFunc("POST /data", HandleData)

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
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	fmt.Println(string(b))
}
