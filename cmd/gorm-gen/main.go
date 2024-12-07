package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func loadEnv() error {
	return godotenv.Load("./cmd/gorm-gen/.env")
}

func main() {
	if err := loadEnv(); err != nil {
		slog.Error("gorm-gen", "Failed to load env file.", err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/infrastructure/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	gormdb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		slog.Error("Failed to create gormdb instance.")
	}
	g.UseDB(gormdb)
	g.ApplyBasic(g.GenerateModel("users"))
	g.Execute()
}
