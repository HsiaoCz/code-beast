package store

import (
	"context"

	"github.com/HsiaoCz/code-beast/some/types"
)

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}
