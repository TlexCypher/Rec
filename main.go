package main

import (
	"Vox/db"
	"fmt"
	"log"
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
	SetupMiddleware(e)
	SetupServer(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
