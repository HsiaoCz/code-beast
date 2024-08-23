package types

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username         string `gorm:"column:username" json:"username"`
	User_ID          string `gorm:"column:user_id" json:"user_id"`
	Password         string `gorm:"column:password" json:"-"`
	Email            string `gorm:"column:email" json:"email"`
	Avatar           string `gorm:"column:avatar" json:"avatar"`
	Background_Image string `gorm:"column:background_image" json:"background_image"`
}

type CreateUserParams struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func (c CreateUserParams) Validate() error {
	if c.Password != c.RePassword {
		return errors.New("check the username or password")
	}
	return nil
}

func NewUserFromParams(params CreateUserParams) *Users {
	return &Users{
		Username:         params.Username,
		Email:            params.Email,
		User_ID:          uuid.New().String(),
		Password:         encryptPassword(params.Password),
		Avatar:           "./picture/avatar/1233.jpg",
		Background_Image: "./picture/bgi/1233.jpg",
	}
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
