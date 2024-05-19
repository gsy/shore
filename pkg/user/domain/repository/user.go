package user_repository

import (
	"context"

	user_model "github.com/gsy/store/pkg/user/domain/model"
)

type UserRepository interface {
	// Create create user
	Create(ctx context.Context, user user_model.User) (err error)
	// FindByID find user by id
	FindByID(ctx context.Context, id uint64) (user user_model.User, err error)
}
