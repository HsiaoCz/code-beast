package data

import (
	"context"

	"github.com/HsiaoCz/code-beast/menlo/types"
)

type SessionDatar interface {
	CreateSession(context.Context, *types.Session) (*types.Session, error)
}
