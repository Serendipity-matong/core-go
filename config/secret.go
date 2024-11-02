package config

import "os"

var (
	DB_ADDR     = os.Getenv("DB_ADDR")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DATABASE = os.Getenv("DB_DATABASE")
)

// token var
var (
	Secret     = "MiMeng"       // 加盐
	ExpireTime = 3600 * 24 * 30 // token有效期
)
