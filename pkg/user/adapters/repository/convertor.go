package repository

import (
	"github.com/gsy/store/pkg/user/domain/user"
	"github.com/gsy/store/pkg/user/adapters/repository/dao/model"
)

type convertor struct{}

func (c convertor) FromEntity(entity *user.User) (obj *model.User) {
	return &model.User{
		UserID:   entity.ID.String(),
		Username: entity.Name.String(),
	}
}

func (c convertor) ToEntity(obj *model.User) (entity *user.User) {
	return &user.User{
		ID:   user.ID(obj.UserID),
		Name: user.Name(obj.Username),
	}
}
