package slick

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Slick struct {
	router       *httprouter.Router
	ErrorHandler ErrorHandler
	middlewares  []Handler
}

func New() *Slick {
	return &Slick{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (s *Slick) Get(path string, h Handler, plugs ...Handler) {
	s.Plugs(plugs...)
	s.router.GET(path, s.makeHttpHandler(h))
}

func (s *Slick) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func (s *Slick) Plugs(middlewares ...Handler) {
	s.middlewares = append(s.middlewares, middlewares...)
}
