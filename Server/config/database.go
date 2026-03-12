package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kerokume-go/schemas"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := NewLogger("DATABASE")

	// Load .env
	erro := godotenv.Load()
	if erro != nil {
		fmt.Errorf("Erro ao carregar .env || %v", erro)
	}

	/// Create and Connect Database
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errf("Postgres opening error: %v", err)
		return nil, err
	}

	// Migrate did schema

	err = db.AutoMigrate(&schemas.Food{})

	if err != nil {
		logger.Errf("Postgres Automigration - Food error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Menu{})

	if err != nil {
		logger.Errf("Postgres Automigration - Menu error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Restaurant{})

	if err != nil {
		logger.Errf("Postgres Automigration - Restaurant error: %v", err)
		return nil, err
	}

	return db, nil
}
