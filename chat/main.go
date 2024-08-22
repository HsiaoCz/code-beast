package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-beast/chat/db"
	"github.com/HsiaoCz/code-beast/chat/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := db.InitDB(); err != nil {
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
		router.HandleFunc("POST /user/login", handlers.TransferHandlerfunc(userHandler.HandleUserLogin))
	}

	http.ListenAndServe(port, router)
}
