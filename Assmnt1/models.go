package main

import (
	"encoding/json"
	"time"
	//"github.com/ha-wk/Assmnt1/database"
)

type Blockstatus string

const (
	Status1  Blockstatus = "pending"
	Status2 Blockstatus = "committed"
)

type Transaction struct {
	ID   int32 `json:"id"`   //SIMNO.  
	Val   int32 `json:"val"`
	Ver   float32 `json:"ver"`
	Valid bool `json:"valid"`
}
type Block_info struct {
	Block_No     int64   `json:"blockNumber"`
	Timestamp    time.Time   `json:"timestamp"`
	TxnsList []Transaction    `json:"transaction"`

	//enum of committed and pending..???
	PrevBlkHash string   `json:"prevBlockHash"`
	Status   Blockstatus  `json:"status"`
}

type Block_methods interface {
	PushValtxns(jsonData []byte,ldb Leveldbinterface) error
	Validate() bool
	UpdateBlkStts()
}


func (block *Block_info)PushValtxns(jsonData []byte,ldb LevelDB) error {
    var txns []Transaction
    err := json.Unmarshal(jsonData, &txns)
    if err != nil {
        return err
    }

    for _, newTxn := range txns {
      existingEntry, err := ldb.Get(string(newTxn.ID)) //modfcs in getentry params
      if err != nil {
        return err
	  }

            if newTxn.Ver == existingEntry.Ver && newTxn.ID == existingEntry.ID {
             // if validateTransaction(string(key),newEntry) {
                 existingEntry.Ver += 1.0
				 existingEntry.Val=newTxn.Val
				 existingEntry.Valid = true 
				 } else {
				existingEntry.Valid = false//marking invalid transaction
			}
		
				err = ldb.Put(string(newTxn.ID), existingEntry)
				if err != nil {
					return err
            }
        }
    
    return nil
}

func (block *Block_info) UpdateBlkStts() {
	
	block.Status = Status2     // Update the status field to "committed"
}

















/*

for _, newTxn := range txns {
        for _, newEntry := range newTxn {
            existingEntry, err := GetEntry(string(newEntry.ID)) //modfcs in getentry params
            if err != nil {
                return err
            }
			*/