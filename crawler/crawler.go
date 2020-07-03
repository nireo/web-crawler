package crawler

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"github.com/nireo/crawler/database"
)

func processAddress(urlString *url.URL, db *gorm.DB, display bool) {
	// check if there is a entry for the host
	url := urlString.String()
	var exists database.Item
	if err := db.Where("url = ?", url).First(&exists).Error; err != nil {
		item := database.Item{
			URL:    urlString.String(),
			Scheme: urlString.Scheme,
			Host:   urlString.Host,
			Path:   urlString.Path,
		}

		// Save to the database
		db.NewRecord(item)
		db.Save(&item)

		// Show what websites have been indexed
		fmt.Println(item.URL + " has been indexed.")
	}
}

// StartCrawler is the crawler part of the application
func StartCrawler(startURL string, db *gorm.DB, display bool) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnRequest(func(r *colly.Request) {
		processAddress(r.URL, db, display)
	})

	c.Visit(startURL)
}
