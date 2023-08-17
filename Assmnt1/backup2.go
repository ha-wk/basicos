

/*package main

import (
	"encoding/json"
	//"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type Leveldbinterface interface {
	//NewLevelDB(dbPath string)(*LevelDB,error)
	Put(key string ,newEntry Transaction)(error)
	Get(key string)(Transaction,error)
}

type LevelDB struct {
	db *leveldb.DB
}

/*func NewLevelDB(dbPath string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}

	return &LevelDB{db: db}, nil
}

func (ldb *LevelDB) Close() {
	ldb.db.Close()
}*/


//var db *leveldb.DB


/*func (ldb *LevelDB) Get(key string) (Transaction, error) {
    data, err := ldb.db.Get([]byte(key), nil)
    if err != nil {
        return Transaction{}, err
    }

    var TxnById Transaction
    err = json.Unmarshal(data, &TxnById)
    if err != nil {
        return Transaction{}, err
    }

    return TxnById, nil
}

func(ldb *LevelDB)Put(key string, newEntry Transaction) error {
    newData, err := json.Marshal(newEntry)
    if err != nil {
        return err
    }

    err = ldb.db.Put([]byte(key), newData, nil)
    if err != nil {
        return err
    }

    return nil
}
*/













/*

func (ldb *LevelDB) Get(key string) (Txn_inp, error) {
	data,err := ldb.db.Get([]byte(key), nil)
	var txndata Txn_inp
	err = json.Unmarshal([]byte(data), &txndata)
	return txndata, nil
}

//getall() left


func (ldb *LevelDB) Put(key, value Txn_inp) error {

	jsonstr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return ldb.db.Put([]byte(key), []byte(jsonstr), nil)

}



func validateTransaction(key string, newTransaction Transaction) bool {
    existingEntry, err := GetEntry(key)
    if err != nil {
        return err
    }

    // Perform your validation logic here
    if newTransaction.Val < existingEntry.Ver {
        return fmt.Errorf("Validation failed: New value is less than existing value")
    }

    return nil
}


*/








/*package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	db, err := leveldb.OpenFile("db", nil)
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
