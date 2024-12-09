package api

import (
	di "Vox/cmd/wire-gen"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"Vox/internal/usecase"
	"Vox/openapi"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func RegisterCategory(ctx echo.Context, categoryName string) error {
	categoryNameVO, err := valueobject.NewCategoryNameFromString(categoryName)
	if err != nil {
		slog.Error("RegisterCategory", "Failed to create categoryName value object.", err)
		return err
	}
	categoryEntity, err := entity.NewCategory(categoryNameVO)
	if err != nil {
		slog.Error("RegisterCategory", "Failed to create category entity.", err)
		return err
	}
	in := usecase.RegisterCategoryInput{
		Category: categoryEntity,
	}
	u := usecase.RegisterCategory{
		CategoryRepository: lo.ToPtr(di.InitCategoryRepository()),
	}
	out, err := u.Do(ctx, in)
	if err != nil {
		slog.Error("RegisterCategory", "Failed to do RegisterCategory usecase.", err)
	}
	return ctx.JSON(http.StatusCreated, openapi.CreateCategoryResponse{
		CategoryId: lo.ToPtr(out.CategoryId),
	})
}
