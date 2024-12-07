package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"

	"github.com/labstack/echo/v4"
)

type GetAllUsersInput struct{}

type GetAllUsersOutput struct {
	Users []entity.User
}

type GetAllUsers struct {
	UserRepository domain.UserRepository
}

func (u *GetAllUsers) Do(ctx echo.Context) (GetAllUsersOutput, error) {
	return GetAllUsersOutput{}, nil
}
