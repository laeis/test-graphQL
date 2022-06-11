package repository

import "gorm.io/gorm"

type User struct {
	connection *gorm.DB
}

func NewUser(c *gorm.DB) *User {
	return &User{
		connection: c,
	}
}
