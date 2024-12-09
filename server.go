package main

import (
	"Vox/internal/api"
	"Vox/openapi"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type VoxSrv struct{}

func (vs *VoxSrv) HealthCheck(ctx echo.Context) error {
	healthCheck := "HealthCheck OK!"
	return ctx.JSON(http.StatusOK,
		openapi.HealthCheckResponse{Message: lo.ToPtr(healthCheck)})
}

func (vs *VoxSrv) GetAllUsers(ctx echo.Context) error {
	return api.GetAllUsers(ctx)
}

func (vs *VoxSrv) CreateNewUser(ctx echo.Context) error {
	var req openapi.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		slog.Error("CreateNewUser", "Failed to bind to openapi.CreateUserRequest.", err)
		return err
	}
	return api.RegisterUser(ctx, *req.Username, *req.Role)
}

func (vs *VoxSrv) CreateNewInventory(ctx echo.Context) error {
	var req openapi.CreateInventoryRequest
	if err := ctx.Bind(&req); err != nil {
		slog.Error("CreateNewInventory", "Failed to bind to openapi.CreateInventoryRequest", err)
		return err
	}
	return api.RegisterInventory(
		ctx,
		req.Categories,
		req.ProductName,
		req.ProductCode,
		req.RemainingQuantity,
		req.Remarks,
	)
}

func (vs *VoxSrv) GetInventoryUserId(ctx echo.Context, userId string) error {
	return api.GetInventoriesForEachUser(ctx, userId)
}

func (vs *VoxSrv) CreateNewCategory(ctx echo.Context) error {
	var req openapi.CreateCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		slog.Error("CreateNewCategory", "Failed to bind to openapi.CreateCategoryRequest", err)
		return err
	}
	return api.RegisterCategory(ctx, req.Name)
}

func (vs *VoxSrv) GetCategory(ctx echo.Context) error {
	return api.GetAllCategories(ctx)
}

func SetupServer(e *echo.Echo) *echo.Echo {
	server := VoxSrv{}
	openapi.RegisterHandlers(e, &server)
	e.Static("/openapi", "openapi")
	return e
}
