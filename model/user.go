package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email    string
	Password string
	Status   string
	Avatar   string
	Phone    string
}
