package handlers

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"

	"math/rand"

	"github.com/HsiaoCz/code-beast/thstar/types"
	"github.com/a-h/templ"
	"github.com/sirupsen/logrus"
)

var StatusCode = &Status{Code: http.StatusOK}

type Status struct {
	Code int
}

type HTTPHandleFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HTTPHandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(math.MaxInt64)
		ctx := context.WithValue(r.Context(), types.CtxRequestIDKey, requestID)
		r = r.WithContext(ctx)
		start := time.Now()
		if err := h(w, r); err != nil {
			defer func() {
				logrus.WithFields(logrus.Fields{
					"requestID":      requestID,
					"method":         r.Method,
					"path":           r.URL.Path,
					"remote address": r.RemoteAddr,
					"error message":  err,
				}).Error("the http server error")
			}()
			if e, ok := err.(APIError); ok {
				StatusCode.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				arr := APIError{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = e.Status
				WriteJSON(w, arr.Status, &arr)
			}
		}
		logrus.WithFields(logrus.Fields{
			"requestID":      requestID,
			"method":         r.Method,
			"code":           StatusCode.Code,
			"path":           r.URL.Path,
			"remote address": r.RemoteAddr,
			"cost":           time.Since(start),
		}).Info("new request coming")
	}
}

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

func Render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}
