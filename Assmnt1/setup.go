package main

import (
	//"fmt"
	//"log"
	//"strconv"

//	"github.com/syndtr/goleveldb/leveldb"
)

func main2() {

	/*db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println("Problem opening Database")
	}
    defer db.Close()

	type LevelDB struct {
		db *leveldb.DB
	}

	for i := 1; i <= 1000; i++ {
		key := fmt.Sprintf("SIM%d", strconv.Itoa(i))
		
		value := fmt.Sprintf(`{"val": %d, "ver": %f,"valid":%v}`, i, 1.0, false)


		err := db.Put([]byte(key), []byte(value), nil)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database populated successfully")

	*/

	/*iter := db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	if err := iter.Error(); err != nil {
		log.Fatal(err)

	
	}*/
}
