package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, nil)
}

// Handle requests to the /items endpoint
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

