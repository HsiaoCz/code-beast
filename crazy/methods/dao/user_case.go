package dao

import (
	"context"

	"github.com/HsiaoCz/code-beast/crazy/methods/types"
	"gorm.io/gorm"
)

type UserCaser interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
}

type UserCase struct {
	db *gorm.DB
}

func UserCaseInit(db *gorm.DB) *UserCase {
	return &UserCase{
		db: db,
	}
}

func (u *UserCase) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}

func (u *UserCase) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	return nil, nil
}
