package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"encoding/json"
)

type DefaultTxn struct {
	SIM map[string] LocalTxnInfo`json:"SIM"`
}

// Handler for the POST Transactions endpoint /create
func PutInitialEntry(c *gin.Context) {

	// Decode the JSON payload from the request body
	var Inp_txn []map[string]LocalTxnInfo
	if err := c.ShouldBindJSON(&Inp_txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.PushValtxns(Inp_txn)

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Insertion Succesfull"})
}

func GetAllBlocks(c *gin.Context){

	data := getAllBlocks()
	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON data")
		return
	}
	c.JSON(http.StatusOK, gin.H{ "message": jsonData})
}

func GetBlockById(c *gin.Context){
	id_str := c.Param("id")
	//fmt.Println(id_str)
	id , _ := strconv.Atoi(id_str)
	data := getBlockById(id)

	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing JSON data")
			return
		}
	c.JSON(http.StatusOK, jsonData)
}

func PrintDB(c *gin.Context){
	db.GetallInCsv()
	c.JSON(http.StatusOK, gin.H{"message": "SUCCESSFULLY PRINTED IN CSV EXTERNAL FILE"})
}















// func resetDBHandler(c *gin.Context){
// 	db.PopulateDB()
// 	c.JSON(http.StatusOK, gin.H{"message": "BACK TO DEFAULT MODE"})
// }