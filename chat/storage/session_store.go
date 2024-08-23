package storage

import (
	"context"

	"github.com/HsiaoCz/code-beast/chat/types"
	"gorm.io/gorm"
)

type SessionStorer interface {
	CreateSession(context.Context, *types.Sessions) (*types.Sessions, error)
}
type SessionStore struct {
	db *gorm.DB
}

func SessionStoreInit(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(ctx context.Context, session *types.Sessions) (*types.Sessions, error) {
	return nil, nil
}
