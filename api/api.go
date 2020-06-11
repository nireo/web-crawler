package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/nireo/crawler/database"
)

var db *gorm.DB

func searchHandler(w http.ResponseWriter, r *http.Request) {
	queries, ok := r.URL.Query()["query"]

	if !ok || len(queries[0]) < 1 {
		http.Error(w, "Query not provided", http.StatusBadRequest)
		return
	}

	query := queries[0]

	var items []database.Item
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

func randomSearchHandler(w http.ResponseWriter, r *http.Request) {
	amount := database.GetItemAmount(db)
	// get 10 random search results
	var indices []int
	for {
		if len(indices) == 10 {
			break
		}

		indices = append(indices, rand.Intn(amount))
	}

	var items []database.Item
	for _, value := range indices {
		var item database.Item
		if err := db.Where("id = ?", value).First(&item).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		items = append(items, item)
	}

	itemToJSON, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(itemToJSON)
}

// RunAPI runs the search engine
func RunAPI(dbParam *gorm.DB) {
	port := "3001"
	http.HandleFunc("/random", randomSearchHandler)
	http.HandleFunc("/search", searchHandler)

	db = dbParam

	fmt.Println("Server running on port " + port)

	http.ListenAndServe(":"+port, nil)
}
