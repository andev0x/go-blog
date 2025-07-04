package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port        string
	DBPath      string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		Port:        os.Getenv("PORT"),
		DBPath:      os.Getenv("DB_PATH"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func ConnectDB(cfg *Config) *gorm.DB {
	if cfg.DatabaseURL != "" {
		db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to PostgreSQL: %v", err)
		}
		return db
	}
	// fallback to SQLite
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to SQLite: %v", err)
	}
	return db
}
