package dto

import (
	"Vox/internal/domain/entity"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func TransferToUserDto(e entity.User) User {
	return User{
		Id:       e.Id.String(),
		Username: e.Username.Value(),
		Role:     e.Role.Value(),
	}
}
