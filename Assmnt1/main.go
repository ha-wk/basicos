package main

import (
	//"encoding/json"
	"fmt"
	//"log"
	"io"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

var ldb *leveldb.DB

func main() {
    
	 dbInstance := Create_Database("db")
	 defer dbInstance.db.Close()



	// Populate the database using the LevelDB instance
	PopulateDB(dbInstance.db)
    http.HandleFunc("/process_transactions", func(w http.ResponseWriter, r *http.Request) {
        jsonData, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusBadRequest)
            return
        }

        //var BlockInst Block_info
		BlockInst := Block_info{}
        err = BlockInst.PushValtxns(jsonData, *dbInstance) // Pass the initialized LevelDB instance
        if err != nil {
            http.Error(w, fmt.Sprintf("Error processing transactions: %v", err), http.StatusInternalServerError)
            return
        }

        BlockInst.UpdateBlkStts()

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Transactions processed successfully"))
    })

    http.ListenAndServe(":8080", nil)
}
