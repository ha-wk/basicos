package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func main() {
	// Connect to LevelDB
	var err error
	db, err = leveldb.OpenFile("db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := SetupRouter()

	log.Fatal(router.Run(":8080"))
	log.Println("Server started on :8080")
}

/* curl -i http://localhost:8080/get?key=myKey

   curl -X PUT -H "Content-Type: application/json" -d '{"key": "myKey", "value": "myValue"}' http://localhost:8080/put

    curl -i http://localhost:8080/getall   */
