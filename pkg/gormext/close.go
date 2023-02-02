package gormext

import (
	"gorm.io/gorm"
)

func Close(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return err
	}
	return database.Close()
}
