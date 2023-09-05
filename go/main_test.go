package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

var testDB *sql.DB

func setupTestDB() {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	testDB = db

	// Initialize the test database schema and data here
	// You might want to run migrations and seed test data
}

func TestMain(m *testing.M) {
	setupTestDB()
	defer testDB.Close()

	// Run the tests
	m.Run()
}

func TestGetTransfers(t *testing.T) {
    // Create a new router
    r := mux.NewRouter()

    // Register the handler function for the /transfers endpoint
    r.HandleFunc("/transfers", getTransfers).Methods("GET")

    // Create a new GET request for the /transfers endpoint
    req, err := http.NewRequest("GET", "/transfers", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder
    rr := httptest.NewRecorder()

    // Serve the request using the router
    r.ServeHTTP(rr, req)

    // Check the HTTP status code of the response
    assert.Equal(t, http.StatusOK, rr.Code, "Expected HTTP status code 200")

    // Add more assertions as needed

    // Print debug information
    fmt.Println("Response:", rr.Body.String()) // Check the response body
}


func TestGetTransfer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/transfers/{id}", getTransfer).Methods("GET")

	// TODO: Replace TRANSFER_ID with an actual transfer ID
	req, err := http.NewRequest("GET", "/transfers/TRANSFER_ID", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected HTTP status code 200")
}

func TestCreateAccount(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", createAccount).Methods("POST")

	accountReq := AccountRequest{
		CustomerID:    "CUSTOMER_ID",
		AccountHolder: "John Doe",
		AccountNumber: "12345678",
		RoutingNumber: "021000021",
	}
	reqBody, _ := json.Marshal(accountReq)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected HTTP status code 200")
}
