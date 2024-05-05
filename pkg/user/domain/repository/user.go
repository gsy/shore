package user_repository

import user_model "github.com/gsy/store/pkg/user/domain/model"

type UserRepository interface {
	CreateUser(user user_model.User) (err error)
}
