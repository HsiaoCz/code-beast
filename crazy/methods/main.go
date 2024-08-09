package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-beast/crazy/methods/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	var (
		port        = os.Getenv("PORT")
		userHandler = &handlers.UserHandlers{}
		router      = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandler.HandleCreateUser))
	}
	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("the http server is running")
	http.ListenAndServe(port, router)
}
