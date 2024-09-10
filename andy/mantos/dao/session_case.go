package dao

import (
	"context"

	"github.com/HsiaoCz/code-beast/crazy/methods/types"
	"gorm.io/gorm"
)

type SessionCaser interface {
	CreateSession(context.Context, *types.Session) (*types.Session, error)
}

type SessionCase struct {
	db *gorm.DB
}

func SessionCaseInit(db *gorm.DB) *SessionCase {
	return &SessionCase{
		db: db,
	}
}

func (s *SessionCase) CreateSession(ctx context.Context, session *types.Session) (*types.Session, error) {
	return nil, nil
}
