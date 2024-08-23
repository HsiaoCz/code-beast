package storage

import (
	"context"

	"github.com/HsiaoCz/code-beast/chat/types"
)

type SessionStorer interface {
	CreateSession(context.Context, *types.Sessions) (*types.Sessions, error)
}
