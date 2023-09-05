package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	
	"github.com/google/uuid" 
	_ "modernc.org/sqlite"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Transfer struct {
	ID             string `json:"id"`
	Timestamp      string `json:"timestamp"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	SourceCustomer string `json:"source_customer"`
	SourceAccount  string `json:"source_account_id"`
	DestCustomer   string `json:"destination_customer"`
	DestAccount    string `json:"destination_account_id"`
}

type AccountRequest struct {
	CustomerID       string `json:"customer_id"`
	AccountHolder    string `json:"account_holder_name"`
	AccountNumber    string `json:"account_number"`
	RoutingNumber    string `json:"routing_number"`
}

type AccountResponse struct {
	ID             string `json:"id"`
	CustomerID     string `json:"customer_id"`
	AccountHolder  string `json:"account_holder_name"`
	AccountNumber  string `json:"account_number"`
	RoutingNumber  string `json:"routing_number"`
}

func isValidRoutingNumber(routingNumber string) bool {
	if len(routingNumber) != 9 {
		return false
	}

	routingDigits := []int{}
	for _, char := range routingNumber {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		routingDigits = append(routingDigits, digit)
	}

	formulaResult := (3*(routingDigits[0]+routingDigits[3]+routingDigits[6]) +
		7*(routingDigits[1]+routingDigits[4]+routingDigits[7]) +
		routingDigits[2]+routingDigits[5]+routingDigits[8]) % 10

	return formulaResult == 0
}

func getTransfers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT t.id, t.timestamp, t.amount, t.status, "+
		"src_cust.first_name || ' ' || src_cust.last_name, src_acc.id, "+
		"dest_cust.first_name || ' ' || dest_cust.last_name, dest_acc.id "+
		"FROM transfers t "+
		"JOIN accounts src_acc ON t.source_account_id = src_acc.id "+
		"JOIN customers src_cust ON src_acc.customer_id = src_cust.id "+
		"JOIN accounts dest_acc ON t.dest_account_id = dest_acc.id "+
		"JOIN customers dest_cust ON dest_acc.customer_id = dest_cust.id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	transfers := []Transfer{}
	for rows.Next() {
		var transfer Transfer
		err := rows.Scan(&transfer.ID, &transfer.Timestamp, &transfer.Amount, &transfer.Status,
			&transfer.SourceCustomer, &transfer.SourceAccount, &transfer.DestCustomer, &transfer.DestAccount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transfers = append(transfers, transfer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfers)
}

func getTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transferID := vars["id"]

	var transfer Transfer
	err := db.QueryRow("SELECT t.id, t.timestamp, t.amount, t.status, "+
		"src_cust.first_name || ' ' || src_cust.last_name, src_acc.id, "+
		"dest_cust.first_name || ' ' || dest_cust.last_name, dest_acc.id "+
		"FROM transfers t "+
		"JOIN accounts src_acc ON t.source_account_id = src_acc.id "+
		"JOIN customers src_cust ON src_acc.customer_id = src_cust.id "+
		"JOIN accounts dest_acc ON t.dest_account_id = dest_acc.id "+
		"JOIN customers dest_cust ON dest_acc.customer_id = dest_cust.id "+
		"WHERE t.id = ?", transferID).
		Scan(&transfer.ID, &transfer.Timestamp, &transfer.Amount, &transfer.Status,
			&transfer.SourceCustomer, &transfer.SourceAccount, &transfer.DestCustomer, &transfer.DestAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfer)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var accountReq AccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValidRoutingNumber(accountReq.RoutingNumber) {
		http.Error(w, "Invalid routing number", http.StatusBadRequest)
		return
	}

	var customerExists bool
	err = db.QueryRow("SELECT COUNT(*) FROM customers WHERE id = ?", accountReq.CustomerID).
		Scan(&customerExists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !customerExists {
		http.Error(w, "Customer does not exist", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
    accountID := uuid.New().String()

	// Insert the data into the database
    result, err := db.Exec("INSERT INTO accounts (id, customer_id, account_holder_name, account_number, routing_number) VALUES (?, ?, ?, ?, ?)",
        accountID, accountReq.CustomerID, accountReq.AccountHolder, accountReq.AccountNumber, accountReq.RoutingNumber)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	newID, _ := result.LastInsertId()
	accountRes := AccountResponse{
		ID:             strconv.FormatInt(newID, 10),
		CustomerID:     accountReq.CustomerID,
		AccountHolder:  accountReq.AccountHolder,
		AccountNumber:  accountReq.AccountNumber,
		RoutingNumber:  accountReq.RoutingNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accountRes)
}

func main() {
	var err error
	db, err = sql.Open("sqlite", "../transfers.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/transfers", getTransfers).Methods("GET")
	r.HandleFunc("/transfers/{id}", getTransfer).Methods("GET")
	r.HandleFunc("/accounts", createAccount).Methods("POST")

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
