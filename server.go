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
		openapi.HealthCheckResponse{Data: lo.ToPtr(healthCheck)})
}

func (vs *VoxSrv) GetAllUsers(ctx echo.Context) error {
	return nil
}

func (vs *VoxSrv) CreateNewUser(ctx echo.Context) error {
	var req openapi.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		slog.Error("Failed to bind to openapi.CreateUserRequest.")
		return err
	}
	return api.RegisterUser(ctx, *req.Role, *req.Username)
}

func SetupServer(e *echo.Echo) *echo.Echo {
	server := VoxSrv{}
	openapi.RegisterHandlers(e, &server)
	e.Static("/openapi", "openapi")
	return e
}
