package command

import (
	"context"
	"github.com/gsy/store/pkg/user/domain/user"
)


type RegisterUser struct {
	Username string
}

type RegisterUserCommandHandler struct {
	repo user.UserRepository
}

func NewRegisterUserCommandHandler(repo user.UserRepository) RegisterUserCommandHandler{
	return RegisterUserCommandHandler{repo: repo}
}


func (h RegisterUserCommandHandler) Handle(ctx context.Context, cmd *RegisterUser) (resp *user.User, err error){
	newUser := user.NewUser(cmd.Username)
	return h.repo.Save(ctx, &newUser)
}
