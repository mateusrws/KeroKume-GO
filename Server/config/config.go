package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db 	   *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize DB 

	db, err = InitializeDatabase()

	if err != nil {
		return fmt.Errorf("error initializing database %v", err)
	}

	return nil
}

func GetDB() *gorm.DB{
	return db
}

func GetLooger(prefix string) *Logger {
	// Initialize Logger
	logger := NewLogger(prefix)
	return logger
}