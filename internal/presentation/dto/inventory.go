package dto

import (
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"

	"github.com/samber/lo"
)

type Inventory struct {
	Id                   string   `json:"id"`
	Categories           []string `json:"category"`
	ProductCode          string   `json:"productCode"`
	ProductName          string   `json:"productName"`
	Remarks              []string `json:"remarks"`
	RemainingQuantity    int64    `json:"remainingQuantity"`
	RepresentProductName string   `json:"representProductName"`
	RepresentProductCode string   `json:"representProductCode"`
}

func TransferToInventoryDto(
	e entity.Inventory,
	representProductName valueobject.RepresentProductName,
	representProductCode valueobject.RepresentProductCode,
) Inventory {
	return Inventory{
		Id: e.Id.String(),
		Categories: lo.Map(e.CategoriesEntity(), func(category entity.Category, index int) string {
			return category.NameVO().Value()
		}),
		ProductCode: e.ProductCodeVO().Value(),
		ProductName: e.ProductNameVO().Value(),
		Remarks: lo.Map(e.RemarksVO(), func(remark valueobject.Remark, index int) string {
			return remark.Value()
		}),
		RemainingQuantity:    e.RemainingQuantity.Value(),
		RepresentProductName: representProductName.Value(),
		RepresentProductCode: representProductCode.Value(),
	}
}
