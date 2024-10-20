package config

import (
	"os"

	"github.com/stringGustavo/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//Check if the DB file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create the database file and directory
		if err := os.MkdirAll("./db", os.ModePerm); err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errorf("sqlite auto migration error: %v", err)
        return nil, err
	}
	return db, nil
}
