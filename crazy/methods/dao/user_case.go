package dao

import (
	"context"

	"github.com/HsiaoCz/code-beast/crazy/methods/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserCaser interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
	DeleteUserByID(context.Context, string) error
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
	tx := u.db.WithContext(ctx).Debug().Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	logrus.WithFields(logrus.Fields{
		"RequestID": ctx.Value(types.CtxRequestIDKey).(int64),
	}).Info("create user request")
	return user, nil
}

func (u *UserCase) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	tx := u.db.WithContext(ctx).Debug().Model(&types.User{}).Find(&user, "user_id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	logrus.WithFields(logrus.Fields{
		"RequestID": ctx.Value(types.CtxRequestIDKey).(int64),
	}).Info("get user by user_id request")
	return &user, nil
}

func (u *UserCase) DeleteUserByID(ctx context.Context, id string) error {
	tx := u.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&types.User{})
	return tx.Error
}
