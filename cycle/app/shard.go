package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

var StatusCode = &Status{Code: http.StatusOK}

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a ErrorMsg) Error() string {
	return a.Message
}

func ErrorMessage(status int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  status,
		Message: message,
	}
}

type Status struct {
	Code int
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(ErrorMsg); ok {
				StatusCode.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				arr := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = arr.Status
				WriteJSON(w, arr.Status, &arr)
			}
		}
		slog.Info("new request coming", "method", r.Method, "code", StatusCode.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}
