package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type RegisterUserInput struct {
	User entity.User
}

type RegisterUserOutput struct {
	Id *valueobject.UserId
}

type RegisterUser struct {
	UserRepository domain.UserRepository
}

func (u *RegisterUser) Do(ctx echo.Context, input RegisterUserInput) (RegisterUserOutput, error) {
	if err := u.UserRepository.Save(input.User); err != nil {
		slog.Error("RegisterUser usecase", "Failed to register user.", err)
		return RegisterUserOutput{
			Id: nil,
		}, err
	}
	return RegisterUserOutput{
		Id: lo.ToPtr(input.User.IdVO()),
	}, nil
}
