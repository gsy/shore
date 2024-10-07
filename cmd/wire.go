//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gsy/store/pkg/common/database"
	"github.com/gsy/store/pkg/user/adapters/repository"
	"github.com/gsy/store/pkg/user/application"
	"github.com/gsy/store/pkg/user/domain/user"
	"github.com/gsy/store/pkg/user/ports"
)


func InitializeUserServer() (ports.HttpServer, error) {
	wire.Build(
		database.GetDatabaseConnection,
		repository.NewRepository,
		wire.Bind(new(user.UserRepository), new(*repository.RepoImpl)),
		application.NewApplication,
		ports.NewHttpServer,
	)
	return ports.HttpServer{}, nil
}
