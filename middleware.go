package main

import (
	"Vox/db"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conn, err := db.Open()
			if err != nil {
				slog.Error("Failed to get database connection with mysql docker instance.")
				return c.String(http.StatusInternalServerError, err.Error())
			}
			config := Config{Conn: conn}
			c.Set("config", config)
			return next(c)
		}
	})
}
