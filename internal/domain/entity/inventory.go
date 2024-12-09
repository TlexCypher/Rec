package entity

import (
	"Vox/internal/domain/valueobject"
	"fmt"
	"log/slog"
)

type Inventory struct {
	Id                valueobject.InventoryId
	Categories        []valueobject.Category
	ProductCode       valueobject.ProductCode
	ProductName       valueobject.ProductName
	Remarks           []valueobject.Remark
	RemainingQuantity valueobject.RemainingQuantity
}

func NewInventory(categories []valueobject.Category,
	productCode valueobject.ProductCode,
	productName valueobject.ProductName,
	remarks []valueobject.Remark,
	remainingQuantity valueobject.RemainingQuantity,
) (Inventory, error) {
	id, err := valueobject.NewInventoryId()
	if err != nil {
		slog.Error("Inventory", "Failed to create new inventory.", err)
		return Inventory{}, fmt.Errorf("Failed to create new inventory.\n")
	}

	return Inventory{
		Id:                id,
		Categories:        categories,
		ProductCode:       productCode,
		ProductName:       productName,
		Remarks:           remarks,
		RemainingQuantity: remainingQuantity,
	}, nil
}

func (e *Inventory) IdVO() valueobject.InventoryId {
	return e.Id
}

func (e *Inventory) CategoriesVO() []valueobject.Category {
	return e.Categories
}

func (e *Inventory) ProductCodeVO() valueobject.ProductCode {
	return e.ProductCode
}

func (e *Inventory) ProductNameVO() valueobject.ProductName {
	return e.ProductName
}

func (e *Inventory) RemarksVO() []valueobject.Remark {
	return e.Remarks
}

func (e *Inventory) RemainingQuantityVO() valueobject.RemainingQuantity {
	return e.RemainingQuantity
}
