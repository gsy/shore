package database

import (
	user_model "github.com/gsy/store/pkg/user/domain/model"
)

//go:generate go run ../../../../main.go gengorm
type UserRepository interface {
	// insert into user (`user_id`, `name`) values (@user.id, @user.name)
	CreateUser(user user_model.User) (err error)
}
