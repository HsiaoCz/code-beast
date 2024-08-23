package storage

import (
	"context"

	"github.com/HsiaoCz/code-beast/chat/types"
)


type UserStorer interface{
	CreateUser(context.Context,*types.Users)(*types.Users,error)
}