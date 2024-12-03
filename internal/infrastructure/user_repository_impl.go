package infrastructure

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type UserRepositoryImpl struct{}

func (ur *UserRepositoryImpl) Save(entity.User) error {
	return nil
}

func (ur *UserRepositoryImpl) Update(entity.User) error {
	return nil
}

func (ur *UserRepositoryImpl) FindById(valueobject.UserId) (*entity.User, error) {
	return nil, nil
}

func (ur *UserRepositoryImpl) DeleteById(valueobject.UserId) error {
	return nil
}
