package repository

import (
	"context"

	repo "github.com/gsy/store/internal/infrastructure/repository"
	"github.com/gsy/store/pkg/user/domain/user"
	"github.com/gsy/store/pkg/user/adapters/repository/dao/model"
	"github.com/gsy/store/pkg/user/adapters/repository/dao/query"
	"gorm.io/gorm"
)

type RepoImpl struct {
	general *repo.GeneralRepository[model.User, user.User]
	dao     *query.Query
	conv    repo.Convertor[model.User, user.User]
}

func NewRepository(db *gorm.DB) *RepoImpl {
	return &RepoImpl{
		general: repo.NewRepository[model.User, user.User](db, convertor{}),
		dao:     query.Use(db),
		conv:    convertor{},
	}
}

func (repo *RepoImpl) FindByID(ctx context.Context, userID string) (entity *user.User, err error) {
	u := repo.dao.User
	obj, err := repo.dao.WithContext(ctx).User.Where(u.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}
	return repo.conv.ToEntity(obj), nil
}


func (repo *RepoImpl) Save(ctx context.Context, newUser *user.User) (result *user.User, err error) {
	return nil, nil
}
