package domain

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type CategoryRepository interface {
	Save(entity.Category) error
	FindById(valueobject.CategoryId) entity.Category
	FindAll() []entity.Category
	Update(entity.Category) error
	DeleteById(valueobject.CategoryId) error
}
