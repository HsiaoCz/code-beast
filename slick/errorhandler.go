package slick

import "log/slog"

type ErrorHandler func(error, *Context) error

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}
