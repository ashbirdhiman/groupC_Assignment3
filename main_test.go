package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var createdItemID string

// TestAddItem tests the AddItem functionlity
func TestAddItem(t *testing.T) {

	item := Item{Name: "Test Item"}

	itemJSON, _ := json.Marshal(item)

	req, err := http.NewRequest("POST", "/AddItem", bytes.NewBuffer(itemJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddItem)

	handler.ServeHTTP(rr, req)

	// check if the status code is what we expect
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Test handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var newItem Item
	err = json.Unmarshal(rr.Body.Bytes(), &newItem)
	if err != nil {
		t.Fatal(err) // if error, stop the test
	}
	// Checking Item name
	if newItem.Name != "Test Item" {
		t.Errorf("handler returned unexpected body: got name %v want %v", newItem.Name, "Test Item")
	}
	// checking if an ID was generated for the new item
	if newItem.ID == "" {
		t.Errorf("Didnt recieved ID for an Item by handler")
	} else {
		// saving the created item ID, which can be used in further test cases by other collaborator.
		createdItemID = newItem.ID
	}
}
