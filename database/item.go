package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Item model for storing websites
type Item struct {
	gorm.Model
	URL    string `json:"url"`
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	Path   string `json:"path"`
}

// GetItemAmount returns the length of all the items in the database
func GetItemAmount(db *gorm.DB) int {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		log.Fatal("error finding items")
	}

	return len(items)
}
