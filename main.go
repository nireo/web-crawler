package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Global db pointer we can use everywhere
var db *gorm.DB

// Item model for storing websites
type Item struct {
	gorm.Model
	URL    string
	Scheme string
	Host   string
	Path   string
}

func processAddress(urlString *url.URL) {
	// check if there is a entry for the host
	var item Item
	if err := db.Where("host = ?", urlString.Host).First(&item).Error; err != nil {
		item := Item{
			URL:    urlString.String(),
			Scheme: urlString.Scheme,
			Host:   urlString.Host,
			Path:   urlString.Path,
		}

		// Save to the database
		db.NewRecord(item)
		db.Save(&item)

		// Show what websites have been indexed
		fmt.Println(urlString.Host + " has been indexed.")
	}
}

func initializeDB() {
	// Open the database file
	db, _ = gorm.Open("sqlite3", "./crawl.db")

	// Migrate the item model to the database
	db.AutoMigrate(&Item{})
	fmt.Println("connected successfully")
}

func main() {
	c := colly.NewCollector()

	// Load up the database
	initializeDB()

	// Take start url from the command line
	var start string
	flag.StringVar(&start, "website", "", "https://github.com")
	flag.Parse()

	// For every link, visit that link
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Print the amount of websites indexed
	fmt.Println(getItemAmount())

	// Process the url information we get from the request
	c.OnRequest(func(r *colly.Request) {
		processAddress(r.URL)
	})

	// Visit the first URL
	c.Visit(start)
}

func getItemAmount() int {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		log.Fatal("error finding items")
	}

	// Just return the length
	return len(items)
}
