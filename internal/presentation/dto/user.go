package dto

import (
	"Vox/internal/domain/entity"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func TransferToDto(e entity.User) User {
	fmt.Println(e)
	return User{
		Username: e.Username.Value(),
		Role:     e.Role.Value(),
	}
}
