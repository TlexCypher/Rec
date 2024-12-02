package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func loadEnv() error {
	return godotenv.Load()
}

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables.")
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
