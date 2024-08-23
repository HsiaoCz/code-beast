package types

import "gorm.io/gorm"

type Sessions struct {
	gorm.Model
	Token      string `gorm:"column:token" json:"token"`
	User_ID    string `gorm:"column:user_id" json:"user_id"`
	IP_Address string `gorm:"column:ip_address" json:"ip_address"`
	User_Agent string `gorm:"column:user_agent" json:"user_agent"`
	Expires_At string `gorm:"column:Expires_at" json:"expires_at"`
}
