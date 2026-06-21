package db

import (
	"fmt"
	"penga/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("./db/sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// Enable foreign keys
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	_, err = sqlDB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	// Auto migrate all models
	err = autoMigrate()
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func autoMigrate() error {
	return DB.AutoMigrate(
		&models.Eier{},
		&models.Inntekt{},
		&models.Utgift{},
		&models.Lån{},
		&models.Fond{},
		&models.Konto{},
		&models.FondSparing{},
		&models.KontoSparing{},
		&models.LånNedbetaling{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
