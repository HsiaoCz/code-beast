package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type ApiError struct {
	Err    error `json:"error"`
	Status int   `json:"status"`
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func TransHTTPHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("error", "err", err)
			WriteJSON(w, http.StatusInternalServerError, ApiError{Err: err, Status: http.StatusInternalServerError})
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func main() {
	http.HandleFunc("/user", TransHTTPHandler(handleGetUserByID))
	http.ListenAndServe(":9001", nil)
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return WriteJSON(w, http.StatusMethodNotAllowed, ApiError{Err: errors.New("method not allowed"), Status: http.StatusMethodNotAllowed})
	}
	user := &User{
		ID:       rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(math.MaxInt64),
		Username: "James",
		Content:  "my man",
	}
	return WriteJSON(w, http.StatusOK, user)
}
