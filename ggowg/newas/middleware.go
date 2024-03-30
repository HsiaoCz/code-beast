package main

import (
	"log/slog"
	"net/http"
)

func RequestLoggerMiddle(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received new request", "method", r.Method, "path", r.URL.Path, "remote-adress", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
