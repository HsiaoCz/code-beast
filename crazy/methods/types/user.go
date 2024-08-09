package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username         string `gorm:"username" json:"username"`
	User_ID          string `gorm:"user_id" json:"user_id"`
	Password         string `gorm:"password" json:"-"`
	Email            string `gorm:"email" json:"email"`
	Avatar           string `gorm:"avatar" json:"avatar"`
	Background_Image string `gorm:"background_image" json:"background_image"`
}
