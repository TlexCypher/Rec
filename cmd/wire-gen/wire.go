//go:build wireinject
// +build wireinject

package wiregen

import (
	"Vox/db"
	"Vox/internal/infrastructure"

	"github.com/google/wire"
)

func InitUserRepositoryImpl() infrastructure.UserRepositoryImpl {
	wire.Build(infrastructure.NewUserRepositoryImpl, db.NewDBConn)
	return infrastructure.UserRepositoryImpl{}
}
