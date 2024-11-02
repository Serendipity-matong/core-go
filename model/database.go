package model

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Type    string
	Content string
}
