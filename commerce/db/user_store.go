package db

import (
	"context"

	"github.com/HsiaoCz/code-beast/commerce/models"
)

type UserStorer interface {
	GetUsers(context.Context) (*models.User, error)
}


