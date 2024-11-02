package model

import "time"

type User struct {
	UserID     string    `gorm:"primary_key;not null" json:"user_id"`
	UserName   string    `gorm:"size:32;not null;unique" json:"username"`
	Password   string    `gorm:"size:64;not null" json:"password"`
	QQ         int       `gorm:"not null;unique" json:"qq"`
	VipDate    time.Time `gorm:"default:null" json:"vip_date"` // 设置默认值为 NULL
	IsAdmin    bool      `gorm:"default:false" json:"is_admin"`
	LastSignIn time.Time `gorm:"default:null" json:"last_sign_in"` // 设置默认值为 NULL
	Points     int       `gorm:"default:0" json:"points"`
}
