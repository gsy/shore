package repository

import (
	"context"

	repo "github.com/gsy/store/internal/infrastructure/repository"
	user_model "github.com/gsy/store/pkg/user/domain/model"
	"github.com/gsy/store/pkg/user/infrastructure/repository/dao/model"
	"github.com/gsy/store/pkg/user/infrastructure/repository/dao/query"
	"gorm.io/gorm"
)

type repoImpl struct {
	general *repo.GeneralRepository[model.User, user_model.User]
	dao     *query.Query
	conv    repo.Convertor[model.User, user_model.User]
}

func NewRepository(db *gorm.DB) *repoImpl {
	return &repoImpl{
		general: repo.NewRepository[model.User, user_model.User](db, convertor{}),
		dao:     query.Use(db),
		conv:    convertor{},
	}
}

func (repo *repoImpl) FindByID(ctx context.Context, userID string) (entity *user_model.User, err error) {
	u := repo.dao.User
	obj, err := repo.dao.WithContext(ctx).User.Where(u.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}
	return repo.conv.ToEntity(obj), nil
}
