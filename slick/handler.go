package slick

import "fmt"

type Handler func(c *Context) error

type Plug func(Handler) Handler

func WithAuth(p Handler) Handler {
	return func(c *Context) error {
		fmt.Println("hello from the middleware")
		return p(c)
	}
}
