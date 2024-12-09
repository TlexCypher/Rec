package usecase

import (
	"Vox/internal/domain"

	"github.com/labstack/echo/v4"
)

type GetAllCategoriesInput struct {
}

type GetAllCategoriesOutput struct {
}

type GetAllCategories struct {
	CategoryRepository domain.CategoryRepository
}

func (u *GetAllCategories) Do(ctx echo.Context, in GetAllCategoriesInput) (GetAllCategoriesOutput, error) {
	return GetAllCategoriesOutput{}, nil
}
