package contact

import "context"

type UserContact interface {
	CreateUser(context.Context)
}
