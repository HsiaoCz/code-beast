package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotfier struct{}

func (s *SimpleAccountNotfier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created", "username", account.Username, "email", account.Email)
	return nil
}

// better account notifier
type BetterAccountNotifier struct{}

func (b *BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created", "username", account.Username, "email", account.Email)
	return nil
}

type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (a *AccountHandler) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the request body", "err", err)
		return
	}
	if err := a.AccountNotifier.NotifyAccountCreated(context.Background(), account); err != nil {
		slog.Error("failed to notify account created", "err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func main() {
	mux := mux.NewRouter()
	accountHandler := &AccountHandler{AccountNotifier: &SimpleAccountNotfier{}}
	mux.HandleFunc("/account", accountHandler.HandleCreateAccount).Methods("POST")
	http.ListenAndServe(":3001", mux)
}
