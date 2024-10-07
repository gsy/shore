package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	UserID uint   `gorm:"user_id"`
	Name   string `gorm:"name"`
}
