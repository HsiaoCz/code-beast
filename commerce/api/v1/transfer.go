package v1

import "net/http"

type Handerfunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerfunc(h Handerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if e, ok := err.(APIError); ok {
				WriteJSON(w, e.Status, e)
				return
			}
			aerr := APIError{
				Status: http.StatusInternalServerError,
				Msg:    err.Error(),
			}
			WriteJSON(w, aerr.Status, aerr)
		}
	}
}
