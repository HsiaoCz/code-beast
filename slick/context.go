package slick

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	req    *http.Request
	resp   http.ResponseWriter
	ctx    context.Context
	params httprouter.Params
}

func (c *Context) Render(comp templ.Component) error {
	return comp.Render(c.ctx, c.resp)
}

func (c *Context) Set(key any, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Context) Get(key string) any {
	return c.ctx.Value(key)
}
