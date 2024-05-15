package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/HsiaoCz/code-beast/chitemp/handlers"
	"github.com/go-chi/chi/v5"
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
		UserHandler = handlers.NewUserHandler()
		logger      = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{}))
		port        = os.Getenv("PORT")
		router      = chi.NewRouter()
		srv         = http.Server{
			Handler:      router,
			Addr:         port,
			ReadTimeout:  time.Millisecond * 1500,
			WriteTimeout: time.Millisecond * 1500,
		}
	)

	slog.SetDefault(logger)

	router.Get("/user/show", handlers.TransferHandlerFunc(UserHandler.HandleShow))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
