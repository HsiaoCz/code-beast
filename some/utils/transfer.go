package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Map map[string]any

type Handler func(w http.ResponseWriter, r *http.Request) error

func TransferHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			slog.Error("server handler error", "err", err)
			WriteJSON(w, http.StatusInternalServerError, Map{
				"Message": err.Error(),
			})
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
