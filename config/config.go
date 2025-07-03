package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port string
	DBPath string
}

func Load() *Config {
	return &Config{
		Port:   os.Getenv("PORT"),
		DBPath: os.Getenv("DB_PATH"),
	}
}

func ConnectDB(cfg *Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
