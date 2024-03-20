package slick

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	req    *http.Request
	resp   http.ResponseWriter
	ctx    context.Context
	params httprouter.Params
}
