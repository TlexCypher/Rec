package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"
	"log"

	"github.com/labstack/echo/v4"
)

type GetAllUsersInput struct{}

type GetAllUsersOutput struct {
	Users *[]entity.User
}

type GetAllUsers struct {
	UserRepository domain.UserRepository
}

func (u *GetAllUsers) Do(ctx echo.Context) (GetAllUsersOutput, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		log.Println("Failed to get all users.")
		return GetAllUsersOutput{}, err
	}
	return GetAllUsersOutput{
		Users: users,
	}, nil
}
