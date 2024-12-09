package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"

	"github.com/labstack/echo/v4"
)

type RegisterCategoryInput struct {
	Category entity.Category
}

type RegisterCategoryOutput struct {
	CategoryId valueobject.CategoryId
}

type RegisterCategory struct {
	CategoryRepository domain.CategoryRepository
}

func (rc *RegisterCategory) Do(ctx echo.Context, input RegisterCategoryInput) (RegisterCategoryOutput, error) {
	return RegisterCategoryOutput{}, nil
}
