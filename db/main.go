package db

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN string
}

type DB struct {
	Config Config
	Conn   *gorm.DB
}

func loadEnv() error {
	return godotenv.Load("./db/.env")
}

func Open() (*DB, error) {
	if err := loadEnv(); err != nil {
		slog.Error("db.loadEnv()", "Failed to load environment variables.", err.Error())
		return nil, err
	}

	config := Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		),
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      config.DSN,
		DefaultStringSize:        256,
		DisableDatetimePrecision: false,
		DontSupportRenameIndex:   false,
		DontSupportRenameColumn:  false,
	}), &gorm.Config{})

	if err != nil {
		slog.Error("db.gorm.Open()", "Failed to connect with mysql db docker instance.")
		return nil, err
	}

	return &DB{
		Config: config,
		Conn:   db,
	}, nil
}
