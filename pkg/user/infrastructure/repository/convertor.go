package repository

import (
	user_model "github.com/gsy/store/pkg/user/domain/model"
	"github.com/gsy/store/pkg/user/infrastructure/repository/dao/model"
)

type convertor struct{}

func (c convertor) FromEntity(entity *user_model.User) (obj *model.User) {
	return &model.User{
		UserID:   entity.ID,
		Username: entity.Name,
	}
}

func (c convertor) ToEntity(obj *model.User) (entity *user_model.User) {
	return &user_model.User{
		ID:   obj.UserID,
		Name: obj.Username,
	}
}
