package model

type UserRegistration struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	QQ       int    `json:"qq"`
}

type LoginCredentials struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
