package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(os.Stdout, p)
	return logger
}
