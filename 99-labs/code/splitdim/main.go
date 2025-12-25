package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"splitdim/pkg/api"
	"splitdim/pkg/db/local"
)

var db api.DataLayer

// TransferHandler is a HTTP handler that implements the money transfer API.
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	//This handler should only accept POST requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var request_body api.Transfer
	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Transfer request arrived with body %v", request_body)
	dbError := db.Transfer(request_body)
	if dbError != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API request failed: %s", dbError)
	}

}

// AccountListHandler is a HTTP handler that returns the current balance of each registered user.
func AccountListHandler(w http.ResponseWriter, r *http.Request) {
	//This handler should only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Accounts request arrived")
	accountsList, err := db.AccountList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	json, err := json.Marshal(accountsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	//This handler should only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Clear request arrived")
	transfers, err := db.Clear()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	json, err := json.Marshal(transfers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// ResetHandler is a HTTP handler that allows to zero out all balances.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	//This handler should only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Reset request arrived")
	err := db.Reset()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	// Set the default logger to a fancier log format.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db = local.NewDataLayer()
	//Registering a static HTTP handler, this serves the static webpage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	//Registering the 4 HTTP endpoints
	http.HandleFunc("/api/transfer", TransferHandler)
	http.HandleFunc("/api/accounts", AccountListHandler)
	http.HandleFunc("/api/clear", ClearHandler)
	http.HandleFunc("/api/reset", ResetHandler)
	log.Println("Server listening on http://:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
