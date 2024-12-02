package main

import (
	"Rec/db"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Config struct {
	Conn *db.DB
}

func loadEnv() error {
	return godotenv.Load()
}

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables.")
	}
	e := echo.New()
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
