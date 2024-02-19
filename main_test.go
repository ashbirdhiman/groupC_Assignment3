package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// Creatd by Mohamed Ayan Khatri - 500226334
// This function is used to test the GetAllItems() functions of main.go
func TestGetAllItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/GetAllItems", nil)
	if err != nil {
		t.Fatal(err)
	}

	// creating a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllItems)
	handler.ServeHTTP(rr, req)
	// checking if the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// checking if the response body is what we expect
	var items []Item
	err = json.Unmarshal(rr.Body.Bytes(), &items)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) == 0 {
		t.Errorf("Expected at least one item in the list")
	}
}

// Creatd by Simrandeep singh - 500229180
// This function is used to test the TestGetOneItem() functions of main.go
func TestGetOneItem(t *testing.T) {

	//Here I am using createdItemID variable which holds value of Item ID
	req, err := http.NewRequest("GET", fmt.Sprintf("/GetOneItem/%s", createdItemID), nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOneItem)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var item Item
	err = json.Unmarshal(rr.Body.Bytes(), &item)
	if err != nil {
		t.Fatal(err)
	}
	if item.ID != createdItemID { // Adjust based on your test setup
		t.Errorf("handler returned wrong item: got ID %v want %v", item.ID, "some-item-id")
	}
}

// test for DeleteOneItem() in main.go
// created by Nikhil Kaushik (500223528)
func TestDeleteOneItem(t *testing.T) {

	testItem := Item{ID: "delete-me", Name: "Delete Test Item"} // assigning value to testItem
	items = append(items, testItem)

	// Create a DELETE request to delete the test item
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/DeleteOneItem/%s", createdItemID), nil)
	if err != nil {
		t.Fatal(err) // If there's an error, I stop the test
	}

	rr := httptest.NewRecorder() // recording the response
	handler := http.HandlerFunc(DeleteOneItem)
	handler.ServeHTTP(rr, req)

	// Check if the http status is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	found := false // Verifying that the item intended was deleted

	for _, item := range items {
		if item.ID == createdItemID {
			found = true
			break
		}
	}

	if found {
		t.Errorf("item was not deleted from the store")
	}
}
