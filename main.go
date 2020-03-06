package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var db *gorm.DB

type Item struct {
	URL    string
	Scheme string
	Host   string
	Path   string
}

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		if strings.Contains(href, "https") {
			parsed, err := url.Parse(href)
			if err != nil {
				log.Fatal(err)
			}

			item := Item{
				URL:    href,
				Scheme: parsed.Scheme,
				Host:   parsed.Host,
				Path:   parsed.Path,
			}

			fmt.Println(item)
		}
	}
}

func initializeDB() {
	db, _ = gorm.Open("./crawl.db", "sqlite")
}

func main() {
	response, err := http.Get("https://github.com")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("a").Each(processElement)
}
