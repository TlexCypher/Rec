package api

import (
	di "Vox/cmd/wire-gen"
	"Vox/internal/usecase"
	"Vox/openapi"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func GetAllUsers(ctx echo.Context) error {
	u := usecase.GetAllUsers{
		UserRepository: lo.ToPtr(di.InitUserRepository()),
	}
	out, err := u.Do(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get all registered users.")
	}
	return ctx.JSON(http.StatusOK, openapi.GetAllUsersResponse{
		Users: &out.Users,
	})
}
