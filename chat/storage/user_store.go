package storage

import (
	"context"

	"github.com/HsiaoCz/code-beast/chat/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByID(context.Context, string) (*types.Users, error)
	DeleteUserByID(context.Context, string) error
}

type UserStore struct {
	db *gorm.DB
}

func UserStoreInit(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) CreateUser(ctx context.Context, user *types.Users) (*types.Users, error) {
	tx := u.db.WithContext(ctx).Debug().Model(&types.Users{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	logrus.WithFields(logrus.Fields{
		"RequestID": ctx.Value(types.CtxRequestIDKey).(int64),
	}).Info("create user request")
	return user, nil
}

func (u *UserStore) GetUserByID(ctx context.Context, id string) (*types.Users, error) {
	var user types.Users
	tx := u.db.WithContext(ctx).Debug().Model(&types.Users{}).Find(&user, "user_id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	logrus.WithFields(logrus.Fields{
		"RequestID": ctx.Value(types.CtxRequestIDKey).(int64),
	}).Info("get user by user_id request")
	return &user, nil
}

func (u *UserStore) DeleteUserByID(ctx context.Context, id string) error {
	tx := u.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&types.Users{})
	return tx.Error
}
