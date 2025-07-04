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
	cfg := &Config{
		Port:        os.Getenv("PORT"),
		DBPath:      os.Getenv("DB_PATH"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	log.Printf("DEBUG: PORT=%s, DB_PATH=%s, DATABASE_URL=%s", cfg.Port, cfg.DBPath, cfg.DatabaseURL)
	if cfg.Port == "" {
		log.Println("WARNING: PORT environment variable is not set. Defaulting to 8080.")
		cfg.Port = "8080"
	}
	return cfg
}

func ConnectDB(cfg *Config) *gorm.DB {
	if cfg.DatabaseURL != "" {
		db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to PostgreSQL: %v", err)
		}
		log.Println("Connected to PostgreSQL database.")
		return db
	}
	if cfg.DBPath != "" {
		db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to SQLite: %v", err)
		}
		log.Println("Connected to SQLite database.")
		return db
	}
	log.Fatal("No database configuration found: set DATABASE_URL for production or DB_PATH for local development.")
	return nil
}
