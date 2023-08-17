package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

type Txn_inp struct {
	ID  int    `json:"id"`
	Val string `json:"value"`
	Ver string `json:"version"`
}

func Run() {
	http.HandleFunc("/compare", handleCompare)
	
    
	serverAddr := "localhost:8080"
	fmt.Println("Server is listening on", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}



func handleCompare(w http.ResponseWriter, r *http.Request) {
	// Parse JSON input into User struct
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming JSON data
	var jsonData []Transaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&jsonData); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
var block Block_info


	err := block.PushValtxns(jsonData, db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing transactions: %v", err), http.StatusInternalServerError)
			return
		}

		// Update block status
		block.UpdateBlkStts()

		// Send response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Transactions processed successfully"))
	}
