package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Global db pointer we can use everywhere
var db *gorm.DB

// Item model for storing websites
type Item struct {
	gorm.Model
	URL    string `json:"url"`
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	Path   string `json:"path"`
}

func processAddress(urlString *url.URL) {
	// check if there is a entry for the host
	url := urlString.String()
	var exists Item
	if err := db.Where("url = ?", url).First(&exists).Error; err != nil {
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
		fmt.Println(item.URL + " has been indexed.")
	}
}

func initializeDB() {
	// Open the database file
	db, _ = gorm.Open("sqlite3", "./crawl.db")

	// Migrate the item model to the database
	db.AutoMigrate(&Item{})
	fmt.Println("connected successfully")
}

type Amount struct {
	Amount int
}

func randomSearchHandler(w http.ResponseWriter, r *http.Request) {
	amount := getItemAmount()
	randomNumber := rand.Intn(amount)

	var item Item
	if err := db.Where("id = ?", randomNumber).First(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	itemToJSON, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(itemToJSON)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	queries, ok := r.URL.Query()["query"]

	if !ok || len(queries[0]) < 1 {
		http.Error(w, "Query not provided", http.StatusBadRequest)
		return
	}

	query := queries[0]

	fmt.Println(query)
	var items []Item
	if err := db.Where("host LIKE ?", "%"+query+"%").Find(&items).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	itemsToJSON, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(itemsToJSON)
}

var globalAmount int

func main() {
	c := colly.NewCollector()

	// Load up the database
	initializeDB()

	// Take start url from the command line
	var start string
	flag.StringVar(&start, "website", "", "https://github.com")
	flag.Parse()

	if start == "" {
		port := "3001"
		http.HandleFunc("/random", randomSearchHandler)
		http.HandleFunc("/search", searchHandler)
		fmt.Println("Server running on port " + port)
		http.ListenAndServe(":"+port, nil)
	} else {
		amount := strconv.Itoa(getItemAmount())
		// startTime := time.Now().Unix()
		fmt.Println(amount + " pages have been indexed.")
		// For every link, visit that link
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(e.Request.AbsoluteURL(link))
		})

		// Process the url information we get from the request
		c.OnRequest(func(r *colly.Request) {
			//if globalAmount >= 100 {
			//	timeEnd := time.Now().Unix()
			//	fmt.Println("Got 100 indexed sites in: " + strconv.Itoa(int(timeEnd-startTime)))
			//	return
			//}
			processAddress(r.URL)
		})

		// Visit the first URL
		c.Visit(start)
	}

}

func getItemAmount() int {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		log.Fatal("error finding items")
	}

	// Just return the length
	return len(items)
}
