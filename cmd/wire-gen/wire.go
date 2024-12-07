// +build wireinject

package di 

import (
	"Vox/db"
	"Vox/internal/infrastructure"

	"github.com/google/wire"
)

func InitUserRepository() infrastructure.UserRepositoryImpl {
	wire.Build(infrastructure.NewUserRepositoryImpl, db.NewDBConn)
	return infrastructure.UserRepositoryImpl{}
}
