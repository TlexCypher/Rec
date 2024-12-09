package entity

import (
	"Vox/internal/domain/valueobject"
	"log/slog"
)

type Category struct {
	Id   valueobject.CategoryId
	Name valueobject.CategoryName
}

func NewCategory(name valueobject.CategoryName) (Category, error) {
	id, err := valueobject.NewCategoryId()
	if err != nil {
		slog.Error("NewCategory", "Failed to create new category entity.", err)
		return Category{}, err
	}
	return Category{
		Id:   id,
		Name: name,
	}, nil
}

func (c *Category) IdVO() valueobject.CategoryId {
	return c.Id
}

func (c *Category) NameVO() valueobject.CategoryName {
	return c.Name
}
