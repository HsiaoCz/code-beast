package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotfier struct{}

func (s *SimpleAccountNotfier) NotifyAccountCreated(ctx context.Context, account Account) error {
	logrus.WithFields(logrus.Fields{
		"username": account.Username,
		"email":    account.Email,
	}).Info("new account created")
	return nil
}

// better account notifier
type BetterAccountNotifier struct{}

func (b *BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	logrus.WithFields(logrus.Fields{
		"username": account.Username,
		"email":    account.Email,
	}).Info("new account created")
	return nil
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (a *AccountHandler) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("failed to decode the request body")
		return
	}
	if err := a.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("failed to notify account created")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	mux := mux.NewRouter()
	accountHandler := &AccountHandler{AccountNotifier: &SimpleAccountNotfier{}}
	mux.HandleFunc("/account", accountHandler.HandleCreateAccount).Methods("POST")
	http.ListenAndServe(":9001", mux)
}
