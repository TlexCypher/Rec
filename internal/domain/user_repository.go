package domain

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type UserRepository interface {
	Save(entity.User) error
	Update(entity.User) error
	FindById(valueobject.UserId) (*entity.User, error)
	FindAll() (*[]entity.User, error)
	DeleteById(valueobject.UserId) error
}
