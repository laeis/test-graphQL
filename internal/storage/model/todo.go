package model

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Text   string
	Done   bool
	UserID uint
	User   User
}
