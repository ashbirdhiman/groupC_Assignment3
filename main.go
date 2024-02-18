package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// item represents a to-do item.
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{} // In-memory storage for items
const Dport = ":8012"

func main() {
	http.HandleFunc("/AddItem", AddItem)
	http.HandleFunc("/GetAllItems", GetAllItems)
	http.HandleFunc("/GetOneItem/", GetOneItem)
	http.HandleFunc("/UpdateItem/", UpdateItem)

	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, nil)
}

// Handle requests to the /AddItem endpoint
// Created by Ashbir - 500228410
func AddItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item.ID = uuid.New().String() // Generate a unique ID for the item
		items = append(items, item)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handle requests to the /GetAllItems endpoint
// Created by Jevica - 500218849
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(items)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//Handle Requests to the /GetOneItem/{itemID} endpoint
//Created By Akash - 500218794

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	//Getting the Item ID from URL
	itemID := strings.TrimPrefix(r.URL.Path, "/GetOneItem/")
	fmt.Print(itemID)
	switch r.Method {
	case "GET":
		for _, item := range items {
			fmt.Print(item.ID)
			if item.ID == itemID {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		//If Item ID is not correct
		http.Error(w, "Item not found", http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Use Get Method Only", http.StatusNotFound)
	}

}

// created by Abhinav Mahajan 500230044
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":

		id := strings.TrimPrefix(r.URL.Path, "/UpdateItem/")

		var updatedItem Item
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Updating the item retrieved through ID
		found := false
		for i, item := range items {
			if item.ID == id {
				updatedItem.ID = id    // checking the ID is unchanged
				items[i] = updatedItem // Update the item in the slice
				found = true
				break
			}
		}
		// displaying error if the ID is not found
		if !found {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedItem)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // dealing with error of method
	}
}
