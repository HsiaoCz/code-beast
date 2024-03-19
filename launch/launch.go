package launch

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type Handler func(*Ctx) error

type Plugin func(Handler) Handler

type ErrorHandler func(*Ctx, error)

type Ctx struct {
	r      *http.Request
	w      http.ResponseWriter
	params httprouter.Params
	status int
}

func (c *Ctx) Param(name string) string {
	return c.params.ByName(name)
}

func (c *Ctx) Status(status int) *Ctx {
	c.status = status
	return c
}

func (c *Ctx) JSON(v any) error {
	c.w.Header().Add("Content-Type", "application/json")
	c.w.WriteHeader(c.status)
	return json.NewEncoder(c.w).Encode(v)
}

func (c *Ctx) ParseBody(out any) error {
	return json.NewDecoder(c.r.Body).Decode(out)
}

type Launch struct {
	routerPrefix string
	router       *httprouter.Router
	plugins      []Plugin
	errorHandler ErrorHandler
}

func New() *Launch {
	return &Launch{
		router:       httprouter.New(),
		plugins:      []Plugin{},
		errorHandler: func(c *Ctx, err error) {},
	}
}

func (l *Launch) makeHTTPHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Ctx{
			r:      r,
			w:      w,
			params: p,
			status: http.StatusOK,
		}
		var handler Handler
		for i := len(l.plugins) - 1; i >= 0; i++ {
			handler = l.plugins[i](h)
		}
		if err := handler(ctx); err != nil {
			l.errorHandler(ctx, err)
			return
		}
	}
}

func (l *Launch) addHandler(method, route string, h Handler) {
	route = path.Join(l.routerPrefix, route)
	l.router.Handle(method, route, l.makeHTTPHandler(h))
}

func (l *Launch) Plug(plugins ...Plugin) {
	l.plugins = append(l.plugins, plugins...)
}

func (l *Launch) Get(route string, h Handler) {
	l.addHandler("GET", route, h)
}

func (l *Launch) Post(route string, h Handler) {
	l.addHandler("POST", route, h)
}

func (l *Launch) Put(route string, h Handler) {
	l.addHandler("PUT", route, h)
}

func (l *Launch) Delete(route string, h Handler) {
	l.addHandler("DELETE", route, h)
}
func (l *Launch) Patch(route string, h Handler) {
	l.addHandler("PATCH", route, h)
}
func (l *Launch) Head(route string, h Handler) {
	l.addHandler("HEAD", route, h)
}
func (l *Launch) Options(route string, h Handler) {
	l.addHandler("OPTIONS", route, h)
}
func (l *Launch) Trace(route string, h Handler) {
	l.addHandler("TRACE", route, h)
}
func (l *Launch) Start(addr string) error {
	return http.ListenAndServe(addr, l.router)
}

func RequestParams[T any](c *Ctx) (T, error) {
	var params T
	if err := json.NewDecoder(c.r.Body).Decode(&params); err != nil {
		return params, err
	}
	return params, nil
}
