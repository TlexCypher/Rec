package dto

import "Vox/internal/domain/entity"

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TransferToCategoryDto(e entity.Category) Category {
	return Category{
		Id:   e.IdVO().String(),
		Name: e.NameVO().Value(),
	}
}
