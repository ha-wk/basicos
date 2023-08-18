package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	//"github.com/ha-wk/Assmnt1/database"
)

type Blockstatus string

const (
	Status1 Blockstatus = "pending"
	Status2 Blockstatus = "committed"
)


//ALL FIELDS RELATED TO LEDGER TRANSACTIONS
type LdgrTxn struct {
	ID    int32   `json:"id"` //SIMNO.
	Val   int32   `json:"val"`
	Ver   float32 `json:"ver"`
	Valid bool    `json:"valid"`
	Hash  string  `json:"hash"`
}

// ALL FIELDS RELATED TO BLOCK
type Block_info struct {
	Block_No  int               `json:"blockNumber"`
	Timestamp time.Time            `json:"timestamp"`
	TxnsList  []map[string]LdgrTxn `json:"TxnsList"`

	
	//enum of committed and pending..???
	PrevBlkHash string      `json:"prevBlockHash"`
	//Status      Blockstatus `json:"status"`
}

// ALL FIELDS RELATED TO EACH TRANSACTION
type TnxFields struct {
	BlockNo        int
	WaitTime       time.Duration
	MaxTnx         int
	PrevBlockHash  string
	TrnxNo         int32
	Blockchan      chan Block_info
	PrintBlockChan chan Block_info
	TnxKV          chan LedgerFields
}

// ALL METHODS DEFINED AS INTERFACES RELATED TO- BLOCK
type Block_methods interface {
	PushValtxns(jsonData []byte, ldb Leveldbinterface) error
	Validate() bool
	UpdateBlkStts()
}


//ROUGHLY INITIALIZING FIRST BLOCK TO GET STARTED
func (T TnxFields) Initialization() {
	BlkDtls.BlockNo = 1
	BlkDtls.WaitTime = 10 * time.Second
	BlkDtls.MaxTnx = 3
	BlkDtls.PrevBlockHash = "  "
	BlkDtls.TrnxNo = 1
	BlkDtls.Blockchan = make(chan Block_info, 1)
	BlkDtls.PrintBlockChan = make(chan Block_info, 1)
	BlkDtls.TnxKV = make(chan LedgerFields, 5)

	Block := Block_info{
		Block_No:    1,
		PrevBlkHash: "",
		TxnsList:    make([]map[string]LdgrTxn, 0),
		Timestamp:   time.Now(),
	}

	BlkDtls.Blockchan <- Block
	fmt.Println(BlkDtls)

}

func (block *Block_info) UpdateBlkStts() {

	//block.Status = Status2 // Update the status field to "committed"
}


/*

for _, newTxn := range txns {
        for _, newEntry := range newTxn {
            existingEntry, err := GetEntry(string(newEntry.ID)) //modfcs in getentry params
            if err != nil {
                return err
            }
*/





func (T TnxFields) AddTransaction() {

	fmt.Println("Insertion in Progress")
	go func() {
		for {

			select {
			//waits for the transaction pair to get
			case trnxPair := <-BlkDtls.TnxKV:
				//wait for the block to come
				block := <-BlkDtls.Blockchan

		
				if len(block.TxnsList) == BlkDtls.MaxTnx {   //CHECK WHETHER THE DEFAULT TXN LIMIT CROSSED?
					block = BlkDtls.AddBlock(block)
				}

				//inserting the transation into the block
				if len(block.TxnsList) <= BlkDtls.MaxTnx {
					trnxPair.Trnx.ID = BlkDtls.TrnxNo            //assgning the transaction number
					BlkDtls.TrnxNo += 1                          //incrementing the transaction number for the next transaction
					if len(block.TxnsList) == 0 {                //alloting the block creating time at the time of the first insertion
						block.Timestamp = time.Now()
					}
					//transaction insertion
					block.TxnsList = append(block.TxnsList, map[string]LdgrTxn{trnxPair.Key: trnxPair.Trnx})
				}
				//push back the block to the channel
				BlkDtls.Blockchan <- block
			
			case <-time.After(BlkDtls.WaitTime):    //IF TIME LIMIT CROSSES
				BlkDtls.AutoWrite()

			}

		}
	}()
}

//Go routine func to automatically export if the default time is completed
func (T TnxFields) AutoWrite() {
	
	if len(BlkDtls.Blockchan) != 0 {
		block := <-BlkDtls.Blockchan
		if len(block.TxnsList) != 0 {
			block = BlkDtls.AddBlock(block)
		}
		BlkDtls.Blockchan <- block
	} 
}


// Final func to add Block to My ledger file
func (T TnxFields) AddBlock(block Block_info) Block_info {

	fmt.Println("Processing Time for Block No.", BlkDtls.BlockNo, "is ", time.Since(block.Timestamp))

	blockByteStream, err := json.Marshal(block)
	if err != nil {
		log.Fatal(err, "error while marshalling block")
	}

	BlkDtls.PrevBlockHash =DeriveHash(string(blockByteStream))
	BlkDtls.PrintBlockChan <- block

	BlkDtls.BlockNo++
	block.Block_No = BlkDtls.BlockNo
	block.PrevBlkHash = BlkDtls.PrevBlockHash
	block.TxnsList = make([]map[string]LdgrTxn, 0)
	return block
}


// A Go Routine func to write ledger to File
func (T TnxFields)WriteToFile() {

	fmt.Println("Writing to File")
	go func() {
		for {

			block := <-T.PrintBlockChan
			byteStream, err := json.Marshal(block)
			if err != nil {
				log.Fatal(err)
			}

			f, err := os.OpenFile("MyLedger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err = f.WriteString(string(byteStream) + "\n"); err != nil {
				panic(err)
			}
		}
	}()
}


