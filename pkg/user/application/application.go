package application

import (
	"github.com/gsy/store/pkg/user/application/command"
	"github.com/gsy/store/pkg/user/domain/user"
)


type Application struct {
	Commands Commands
	Queries Queries
}

type Commands struct{
	RegisterUser command.RegisterUserCommandHandler
	LoginUser command.LoginUserCommandHandler
}

type Queries struct{}


func NewApplication(repo user.UserRepository) Application {
	return Application{
		Commands: Commands{
			RegisterUser: command.NewRegisterUserCommandHandler(repo),
			LoginUser: command.NewLoginUserCommandHandler(),
		},
	}
}
