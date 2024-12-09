package api

import (
	di "Vox/cmd/wire-gen"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"Vox/internal/presentation/dto"
	"Vox/internal/usecase"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func RegisterInventory(
	ctx echo.Context, categories []string,
	productName string, productCode string,
	remainingQuantity int, remarks *[]string,
) error {
	categoryEntities := make([]entity.Category, 0)
	for _, category := range categories {
		categoryNameVO, err := valueobject.NewCategoryNameFromString(category)
		if err != nil {
			slog.Error("RegisterInventory", "Failed to generate category name value object.", err)
			return err
		}
		categoryEntity, err := entity.NewCategory(categoryNameVO)
		if err != nil {
			slog.Error("RegisterInventory", "Failed to generate category entity.", err)
			return err
		}
		categoryEntities = append(categoryEntities, categoryEntity)
	}

	productNameVO, err := valueobject.NewProductNameFromString(productName)
	if err != nil {
		slog.Error("RegisterInventory", "Failed to generate productname value object", err)
		return err
	}

	productCodeVO, err := valueobject.NewProductCodeFromString(productCode)
	if err != nil {
		slog.Error("RegisterInventory", "Failed to generate product code value object", err)
		return err
	}

	remainingQuantityVO, err := valueobject.NewRemainingQuantityFromInt64(int64(remainingQuantity))
	if err != nil {
		slog.Error("RegisterInventory", "Failed to generate product remaining quantity value object.", err)
		return err
	}

	remarksVO := make([]valueobject.Remark, 0)
	if remarks != nil {
		for _, remark := range *remarks {
			remarkVO, err := valueobject.NewRemarkFromString(remark)
			if err != nil {
				slog.Error("RegisterInventory", "Failed to generate remark valueobject.", err)
				return err
			}
			remarksVO = append(remarksVO, remarkVO)
		}
	}

	in := usecase.RegisterInventoryInput{
		Categories:        categoryEntities,
		ProductName:       productNameVO,
		ProductCode:       productCodeVO,
		RemainingQuantity: remainingQuantityVO,
		Remarks:           lo.ToPtr(remarksVO),
	}
	u := usecase.RegisterInventory{
		InventoryRepository: lo.ToPtr(di.InitInventoryRepository()),
	}
	out, err := u.Do(ctx, in)
	if err != nil {
		slog.Error("RegisterInventory", "Failed to do RegisterInventory usecase.", err)
		return err
	}
	/*もうすでにentityとして成立しているのでerrを握りつぶしてもOK*/
	representProductName := lo.Must(valueobject.NewRepresentProductNameFromString(out.Inventory.ProductName.Value()))
	representProductCode := lo.Must(valueobject.NewRepresentProductCodeFromString(out.Inventory.ProductCode.Value()))
	/*新しくマスターに登録する時は表示名も同じ*/
	return ctx.JSON(http.StatusCreated,
		dto.TransferToInventoryDto(
			out.Inventory,
			representProductName,
			representProductCode,
		))
}
