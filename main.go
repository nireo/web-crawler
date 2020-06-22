package main

import (
	"flag"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nireo/crawler/api"
	"github.com/nireo/crawler/crawler"
	"github.com/nireo/crawler/database"
)

func main() {
	// Take start url from the command line
	var start string
	flag.StringVar(&start, "website", "", "https://github.com")
	flag.Parse()

	db, err := database.Initialize()
	if err != nil {
		fmt.Printf("Database not found, err: %s", err.Error())
		return
	}

	if start == "" {
		api.RunAPI(db)
	} else {
		crawler.StartCrawler(start, db)
	}
}
