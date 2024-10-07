package command

import (
	"context"
	"fmt"

)


type LoginUser struct {
	Username string
}

type LoginUserCommandHandler struct {}

func NewLoginUserCommandHandler() LoginUserCommandHandler{
	return LoginUserCommandHandler{}
}


func (h LoginUserCommandHandler) Handle(ctx context.Context, cmd *LoginUser) (err error){
	fmt.Printf("user: %q login", cmd.Username)
	return nil
}
