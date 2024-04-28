package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandleFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HTTPHandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(APIError); ok {
				WriteJSON(w, e.Status, &e)
			} else {
				arr := APIError{
					Status: http.StatusInternalServerError,
					Msg:    err.Error(),
				}
				WriteJSON(w, arr.Status, &arr)
			}
		}
		slog.Info("new request comming", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func Render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}
