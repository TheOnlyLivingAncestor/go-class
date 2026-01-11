package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"splitdim/pkg/api"
	"splitdim/pkg/db/kvstore"
	"splitdim/pkg/db/local"
	"syscall"
	"time"
)

var db api.DataLayer

// KVStoreMode defines the data layer mode (local/redis/kvstore).
var KVStoreMode = "local"

// KVStoreAddr stores the key-value store address as a DNS domain name or IP address.
var KVStoreAddr = "localhost:8001"

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
		log.Printf("Transfer request body decoding failed with error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Transfer request arrived with body %v", request_body)
	dbError := db.Transfer(request_body)
	if dbError != nil {
		log.Printf("Database error during Transfer request  %v", dbError)
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
		log.Printf("AccountList retrieval failed with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	json, err := json.Marshal(accountsList)
	if err != nil {
		log.Printf("AccountList marshal failed with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		log.Printf("Failed to write response with error %v", err)
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
		log.Printf("Failed to retrieve required transfers from database with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	json, err := json.Marshal(transfers)
	if err != nil {
		log.Printf("Failed to marshal response with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		log.Printf("Failed to write response with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	//This handler should only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	json, err := json.Marshal("OK")
	if err != nil {
		log.Printf("Failed to marshal response with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "API request failed: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		log.Printf("Failed to write response with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		log.Printf("Failed to reset database with error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func startServer(s *http.Server) {
	log.Println("Server listening on http://:8080")
	err := s.ListenAndServe()
	// http.ErrServerClosed should not be logged to the user
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error %v", err)
	}
}

func main() {
	// Set the default logger to a fancier log format.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Setting up flags
	var kvStoreModeFlag string
	var kvStoreAddrFlag string
	flag.StringVar(&kvStoreModeFlag, "mode", "local", "Specifies which storage layer to use")
	flag.StringVar(&kvStoreAddrFlag, "addr", "localhost:8081", "Specifies the address of storage layer, only relevant if key-value storage is chosen")

	//Parsing the flags
	flag.Parse()

	if os.Getenv("KVSTORE_MODE") != "" {
		KVStoreMode = os.Getenv("KVSTORE_MODE")
	} else {
		KVStoreMode = kvStoreModeFlag
	}
	if os.Getenv("KVSTORE_ADDR") != "" {
		KVStoreAddr = os.Getenv("KVSTORE_ADDR")
	} else {
		KVStoreAddr = kvStoreAddrFlag
	}
	switch KVStoreMode {
	case "kvstore":
		log.Printf("Using the kvstore datalayer at %q", KVStoreAddr)
		db = kvstore.NewDataLayer(KVStoreAddr)
	case "local":
		fallthrough
	default:
		log.Println("Using the local datalayer")
		db = local.NewDataLayer()

	}
	//Registering a static HTTP handler, this serves the static webpage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	//Registering the 4 HTTP endpoints
	http.HandleFunc("/api/transfer", TransferHandler)
	http.HandleFunc("/api/accounts", AccountListHandler)
	http.HandleFunc("/api/clear", ClearHandler)
	http.HandleFunc("/api/reset", ResetHandler)
	http.HandleFunc("/healthz", HealthzHandler)
	//The already existing server should be moved to a goroutine for the graceful shutdown
	s := &http.Server{Addr: ":8080"}
	// Start the server in the goroutine
	go startServer(s)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutdown signal received")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Graceful server shutdown failed with error %v", err)
	} else {
		log.Println("Graceful server sutdown succeeded")
	}

}
