package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"

	"github.com/labstack/echo/v4"
)

type GetInventoryForEachUserInput struct {
	UserId valueobject.UserId
}

type GetInventoryForEachUserOutput struct {
	Inventories          []entity.Inventory
	RepresentProductName valueobject.RepresentProductName
	RepresentProductCode valueobject.RepresentProductCode
}

type GetInventoryForEachUser struct {
	UserRepository      domain.UserRepository
	InventoryRepository domain.InventoryRepository
}

func (u *GetInventoryForEachUser) Do(ctx echo.Context, input GetInventoryForEachUserInput) (GetInventoryForEachUserOutput, error) {
	return GetInventoryForEachUserOutput{}, nil
}
