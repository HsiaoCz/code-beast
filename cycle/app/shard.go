package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type APIError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (a APIError) Error() string {
	return a.Msg
}

type Status struct {
	Code int
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := &Status{Code: http.StatusOK}
		if err := h(w, r); err != nil {
			slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(APIError); ok {
				status.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				arr := APIError{
					Status: http.StatusInternalServerError,
					Msg:    err.Error(),
				}
				status.Code = arr.Status
				WriteJSON(w, arr.Status, &arr)
			}
		}
		slog.Info("new request coming", "method", r.Method, "code", status.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}
