package api

import (
	di "Vox/cmd/wire-gen"
	"Vox/internal/domain/entity"
	"Vox/internal/domain/valueobject"
	"Vox/internal/presentation/dto"
	"Vox/internal/usecase"
	"Vox/openapi"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func GetInventoriesForEachUser(ctx echo.Context, userId string) error {
	userIdVO, err := valueobject.NewUserIdFromString(userId)
	if err != nil {
		slog.Error("GetInventoryForEachUser", "Failed to get inventories for each user.", err)
		return err
	}
	u := usecase.GetInventoryForEachUser{
		UserRepository:      lo.ToPtr(di.InitUserRepository()),
		InventoryRepository: lo.ToPtr(di.InitInventoryRepository()),
	}
	input := usecase.GetInventoryForEachUserInput{
		UserId: userIdVO,
	}
	out, err := u.Do(ctx, input)
	if err != nil {
		slog.Error("GetInventoryForEachUser", "Failed to do usecase.", err)
		return err
	}
	return ctx.JSON(http.StatusOK, openapi.GetInventoriesResponse{
		Inventories: lo.ToPtr(
			lo.Map(out.Inventories, func(inventory entity.Inventory, index int) dto.Inventory {
				return dto.TransferToInventoryDto(
					inventory,
					out.RepresentProductName,
					out.RepresentProductCode,
				)
			}),
		),
	})
}
