package main

import (
	"flag"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nireo/web-crawler/api"
	"github.com/nireo/web-crawler/crawler"
	"github.com/nireo/web-crawler/database"
)

func main() {
	// Take start url from the command line
	var start string
	flag.StringVar(&start, "website", "", "The starting address where the web crawler is planted")

	var display bool
	flag.BoolVar(&display, "display", true, "Display the website url where the web crawler goes to; default: true")
	flag.Parse()

	db, err := database.Initialize()
	if err != nil {
		fmt.Printf("Database not found, err: %s", err.Error())
		return
	}

	if start == "" {
		api.RunAPI(db)
	} else {
		crawler.StartCrawler(start, db, display)
	}
}
