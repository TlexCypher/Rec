package entity

import (
	"Vox/internal/domain/valueobject"
	"fmt"
	"log/slog"
)

type User struct {
	id       valueobject.UserId
	role     valueobject.Role
	username valueobject.Username
}

func NewUser(role valueobject.Role, username valueobject.Username) (User, error) {
	id, err := valueobject.NewUserId()
	if err != nil {
		slog.Error("User", "Failed to create new user.", err)
		return User{}, fmt.Errorf("Failed to create new user.")
	}
	return User{
		id:       id,
		role:     role,
		username: username,
	}, nil
}

func (e *User) IdVO() valueobject.UserId {
	return e.id
}

func (e *User) RoleVO() valueobject.Role {
	return e.role
}

func (e *User) SetRoleVO(newRole valueobject.Role) {
	e.role = newRole
}

func (e *User) UsernameVO() valueobject.Username {
	return e.username
}

func (e *User) SetUsernameVO(newUsername valueobject.Username) {
	e.username = newUsername
}
