package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-beast/crazy/methods/dao"
	"github.com/HsiaoCz/code-beast/crazy/methods/db"
	"github.com/HsiaoCz/code-beast/crazy/methods/handlers"
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
		userCase    = dao.UserCaseInit(db.Get())
		sessionCase = dao.SessionCaseInit(db.Get())
		userHandler = handlers.UserHandlersInit(userCase, sessionCase)
		router      = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandler.HandleCreateUser))
		router.HandleFunc("GET /user", handlers.TransferHandlerfunc(userHandler.HandleGetUserByID))
		router.HandleFunc("DELETE /user", handlers.TransferHandlerfunc(userHandler.HandleDeleteUserByID))
	}
	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("the http server is running")
	http.ListenAndServe(port, router)
}
