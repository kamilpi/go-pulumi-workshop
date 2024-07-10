package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items  = []Item{}
	nextID = 1
	mutex  = &sync.Mutex{}
)

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	mutex.Lock()
	item.ID = nextID
	nextID++
	items = append(items, item)
	mutex.Unlock()
	json.NewEncoder(w).Encode(item)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createItem(w, r)
		} else if r.Method == http.MethodGet {
			getItems(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
}
