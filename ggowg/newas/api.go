package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (a *APIServer) Run() error {
	router := http.NewServeMux()
	userhandler := NewUserHandler()
	router.HandleFunc("POST /user/create", TransferHTTPHandler(userhandler.HandlerCreateUser))

	server := http.Server{
		Addr:    a.addr,
		Handler: RequestLoggerMiddle(router),
	}
	slog.Info("the server is running", "port", a.addr)
	return server.ListenAndServe()
}

func TransferHTTPHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("the internal error", "err", err)
			WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
