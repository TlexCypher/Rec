package usecase

import (
	"Vox/internal/domain"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type RegisterInventoryInput struct {
	Categories        []entity.Category
	ProductName       valueobject.ProductName
	ProductCode       valueobject.ProductCode
	RemainingQuantity valueobject.RemainingQuantity
	Remarks           *[]valueobject.Remark
}

type RegisterInventoryOutput struct {
	Inventory entity.Inventory
}

type RegisterInventory struct {
	InventoryRepository domain.InventoryRepository
}

func (u *RegisterInventory) Do(ctx echo.Context, input RegisterInventoryInput) (RegisterInventoryOutput, error) {
	e, err := entity.NewInventory(input.Categories, input.ProductCode, input.ProductName, *input.Remarks, input.RemainingQuantity)
	if err != nil {
		slog.Error("RegisterInventory.Do", "Failed to create inventory entity.", err)
	}
	if err := u.InventoryRepository.Save(e); err != nil {
		return RegisterInventoryOutput{}, err
	}
	return RegisterInventoryOutput{
		Inventory: e,
	}, nil
}
