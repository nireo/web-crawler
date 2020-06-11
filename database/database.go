package database

import (
	"github.com/jinzhu/gorm"
)

// Initialize the database
func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "../crawl.db")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Item{})
	return db, err
}
