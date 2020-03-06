package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		if strings.Contains(href, "https") {
			fmt.Println(href)
		}
	}
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
