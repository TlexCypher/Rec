package infrastructure

import (
	"Vox/db"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
)

type CategoryRepositoryImpl struct {
	db *db.DB
}

func NewCategoryRepositoryImpl(db *db.DB) CategoryRepositoryImpl {
	return CategoryRepositoryImpl{
		db: db,
	}
}

func (cr *CategoryRepositoryImpl) Save(e entity.Category) error {
	return nil
}

func (cr *CategoryRepositoryImpl) FindById(categoryId valueobject.CategoryId) entity.Category {
	return entity.Category{}
}

func (cr *CategoryRepositoryImpl) FindAll() []entity.Category {
	return []entity.Category{}
}

func (cr *CategoryRepositoryImpl) Update(e entity.Category) error {
	return nil
}

func (cr *CategoryRepositoryImpl) DeleteById(categoryId valueobject.CategoryId) error {
	return nil
}
