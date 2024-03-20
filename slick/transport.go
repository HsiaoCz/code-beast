package slick

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Slick) makeHttpHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Context{
			req:    r,
			resp:   w,
			params: p,
			ctx:    r.Context(),
		}
		if err := h(ctx); err != nil {
			// handle the err
			s.ErrorHandler(err, ctx)
		}
	}
}
