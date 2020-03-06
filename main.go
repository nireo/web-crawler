package main

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Item struct {
	URL    string
	Scheme string
	Host   string
	Path   string
}

func processAddress(urlString *url.URL) {
	item := Item{
		URL:    urlString.String(),
		Scheme: urlString.Scheme,
		Host:   urlString.Host,
		Path:   urlString.Path,
	}

	fmt.Println(item)
}

func initializeDB() {
	db, _ = gorm.Open("./crawl.db", "sqlite")
}

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		processAddress(r.URL)
	})

	c.Visit("http://go-colly.org/")
}
