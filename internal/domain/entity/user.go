package entity

import (
	"Vox/internal/domain/valueobject"
	"fmt"
	"log/slog"
)

type User struct {
	Id       valueobject.UserId
	Role     valueobject.Role
	Username valueobject.Username
}

func NewUser(role valueobject.Role, username valueobject.Username) (User, error) {
	id, err := valueobject.NewUserId()
	if err != nil {
		slog.Error("User", "Failed to create new user.", err)
		return User{}, fmt.Errorf("Failed to create new user.\n")
	}
	return User{
		Id:       id,
		Role:     role,
		Username: username,
	}, nil
}

func (e *User) IdVO() valueobject.UserId {
	return e.Id
}

func (e *User) RoleVO() valueobject.Role {
	return e.Role
}

func (e *User) SetRoleVO(newRole valueobject.Role) {
	e.Role = newRole
}

func (e *User) UsernameVO() valueobject.Username {
	return e.Username
}

func (e *User) SetUsernameVO(newUsername valueobject.Username) {
	e.Username = newUsername
}
